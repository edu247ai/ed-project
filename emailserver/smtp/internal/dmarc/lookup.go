package dmarc

import (
	"context"
	"net"
	"strings"
)

func Lookup(domain string) (string, error) {
	return LookupContext(context.Background(), domain)
}

func LookupContext(ctx context.Context, domain string) (string, error) {
	txts, err := net.DefaultResolver.LookupTXT(ctx, "_dmarc."+domain)
	if err != nil {
		return "", err
	}

	for _, t := range txts {
		if strings.HasPrefix(t, "v=DMARC1") {
			return t, nil
		}
	}

	return "", nil
}
