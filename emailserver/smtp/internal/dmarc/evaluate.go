package dmarc

import (
	"context"
	"errors"
	"net"
	"strings"
)

func Evaluate(spfResult, dkimResult, domain string) (string, error) {
	record, err := Lookup(domain)
	if err != nil {
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) && dnsErr.IsNotFound {
			return "none", nil
		}
		return "", err
	}
	if record == "" {
		return "none", nil
	}
	if spfResult == "pass" || dkimResult == "pass" {
		return "pass", nil
	}
	return Parse(record).Mode, nil
}

func EvaluateAligned(spfResult, spfDomain, dkimResult, dkimDomain, fromDomain string) (string, error) {
	return EvaluateAlignedContext(context.Background(), spfResult, spfDomain, dkimResult, dkimDomain, fromDomain)
}

func EvaluateAlignedContext(ctx context.Context, spfResult, spfDomain, dkimResult, dkimDomain, fromDomain string) (string, error) {
	fromDomain = normalizeDomain(fromDomain)
	if fromDomain == "" {
		return "none", nil
	}
	if spfResult == "pass" && normalizeDomain(spfDomain) == fromDomain {
		return "pass", nil
	}
	if dkimResult == "pass" && normalizeDomain(dkimDomain) == fromDomain {
		return "pass", nil
	}
	record, err := LookupContext(ctx, fromDomain)
	if err != nil {
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) && dnsErr.IsNotFound {
			return "none", nil
		}
		return "", err
	}
	if record == "" {
		return "none", nil
	}
	return Parse(record).Mode, nil
}

func normalizeDomain(domain string) string {
	return strings.TrimSuffix(strings.ToLower(strings.TrimSpace(domain)), ".")
}
