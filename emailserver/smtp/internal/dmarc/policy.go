package dmarc

import "strings"

type Policy struct {
	Mode string
}

func Parse(dmarc string) Policy {
	parts := strings.Split(dmarc, ";")
	p := Policy{Mode: "none"}

	for _, part := range parts {
		part = strings.TrimSpace(part)
		lower := strings.ToLower(part)
		if strings.HasPrefix(lower, "p=") {
			mode := strings.TrimSpace(lower[2:])
			if mode == "none" || mode == "quarantine" || mode == "reject" {
				p.Mode = mode
			}
		}
	}

	return p
}
