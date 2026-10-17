package spf

import (
	"context"
	"net"
	"strings"
)

func Lookup(domain string) (string, error) {
	return LookupContext(context.Background(), domain)
}

func LookupContext(ctx context.Context, domain string) (string, error) {
	txts, err := net.DefaultResolver.LookupTXT(ctx, domain)
	if err != nil {
		return "", err
	}

	for _, t := range txts {
		if strings.HasPrefix(strings.ToLower(t), "v=spf1") {
			return t, nil
		}
	}

	return "", nil
}
