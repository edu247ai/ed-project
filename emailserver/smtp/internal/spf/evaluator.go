package spf

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/netip"
	"strings"
)

func Evaluate(ip, domain string) (string, error) {
	return EvaluateContext(context.Background(), ip, domain)
}

func EvaluateContext(ctx context.Context, ip, domain string) (string, error) {
	clientIP, err := netip.ParseAddr(ip)
	if err != nil {
		return "neutral", fmt.Errorf("invalid SPF client IP %q: %w", ip, err)
	}
	lookups := 0
	return evaluate(ctx, clientIP.Unmap(), strings.TrimSuffix(strings.ToLower(domain), "."), 0, &lookups)
}

func evaluate(ctx context.Context, ip netip.Addr, domain string, depth int, lookups *int) (string, error) {
	if depth > 10 {
		return "permerror", errors.New("SPF include recursion limit exceeded")
	}
	record, err := LookupContext(ctx, domain)
	if err != nil {
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) && dnsErr.IsNotFound {
			return "neutral", nil
		}
		return "neutral", err
	}
	if record == "" {
		return "neutral", nil
	}
	redirect := ""
	for _, mechanism := range Parse(record) {
		if mechanism.Name == "redirect" {
			redirect = mechanism.Arg
			continue
		}
		matched, err := matches(ctx, ip, domain, mechanism, depth, lookups)
		if err != nil {
			return "neutral", err
		}
		if matched {
			switch mechanism.Qualifier {
			case '-':
				return "fail", nil
			case '~':
				return "softfail", nil
			case '?':
				return "neutral", nil
			default:
				return "pass", nil
			}
		}
	}
	if redirect != "" {
		*lookups++
		if *lookups > 10 {
			return "permerror", errors.New("SPF DNS lookup limit exceeded")
		}
		return evaluate(ctx, ip, strings.TrimSuffix(strings.ToLower(redirect), "."), depth+1, lookups)
	}
	return "neutral", nil
}

func matches(ctx context.Context, ip netip.Addr, domain string, mechanism Mechanism, depth int, lookups *int) (bool, error) {
	switch mechanism.Name {
	case "all":
		return true, nil
	case "ip4", "ip6":
		prefix, err := netip.ParsePrefix(mechanism.Arg)
		if err != nil {
			address, parseErr := netip.ParseAddr(mechanism.Arg)
			if parseErr != nil {
				return false, nil
			}
			prefix = netip.PrefixFrom(address, address.BitLen())
		}
		return prefix.Contains(ip), nil
	case "include":
		*lookups = *lookups + 1
		if *lookups > 10 {
			return false, errors.New("SPF DNS lookup limit exceeded")
		}
		result, err := evaluate(ctx, ip, strings.TrimSuffix(strings.ToLower(mechanism.Arg), "."), depth+1, lookups)
		if err != nil {
			return false, err
		}
		return result == "pass", nil
	case "a":
		*lookups = *lookups + 1
		if *lookups > 10 {
			return false, errors.New("SPF DNS lookup limit exceeded")
		}
		target := mechanism.Arg
		if target == "" {
			target = domain
		}
		return domainHasIP(ctx, target, ip)
	case "mx":
		*lookups = *lookups + 1
		if *lookups > 10 {
			return false, errors.New("SPF DNS lookup limit exceeded")
		}
		target := mechanism.Arg
		if target == "" {
			target = domain
		}
		mxs, err := net.DefaultResolver.LookupMX(ctx, target)
		if err != nil {
			var dnsErr *net.DNSError
			if errors.As(err, &dnsErr) && dnsErr.IsNotFound {
				return false, nil
			}
			return false, err
		}
		for _, mx := range mxs {
			found, err := domainHasIP(ctx, strings.TrimSuffix(mx.Host, "."), ip)
			if err != nil || found {
				return found, err
			}
		}
	}
	return false, nil
}

func domainHasIP(ctx context.Context, domain string, ip netip.Addr) (bool, error) {
	addresses, err := net.DefaultResolver.LookupIP(ctx, "ip", domain)
	if err != nil {
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) && dnsErr.IsNotFound {
			return false, nil
		}
		return false, err
	}
	for _, address := range addresses {
		parsed, ok := netip.AddrFromSlice(address)
		if ok && parsed.Unmap() == ip {
			return true, nil
		}
	}
	return false, nil
}
