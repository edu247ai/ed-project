package spf

import "strings"

type Mechanism struct {
	Name      string
	Arg       string
	Qualifier byte
}

func Parse(record string) []Mechanism {
	parts := strings.Fields(record)
	if len(parts) == 0 || !strings.EqualFold(parts[0], "v=spf1") {
		return nil
	}
	mechanisms := make([]Mechanism, 0, len(parts)-1)
	for _, part := range parts[1:] {
		if strings.Contains(part, "=") && !strings.Contains(part, ":") {
			name, arg, _ := strings.Cut(part, "=")
			mechanisms = append(mechanisms, Mechanism{Name: strings.ToLower(name), Arg: arg})
			continue
		}
		qualifier := byte('+')
		if strings.ContainsRune("+-~?", rune(part[0])) {
			qualifier = part[0]
			part = part[1:]
		}
		name, arg, found := strings.Cut(part, ":")
		if !found {
			name, arg, _ = strings.Cut(part, "/")
		}
		mechanisms = append(mechanisms, Mechanism{
			Name:      strings.ToLower(name),
			Arg:       arg,
			Qualifier: qualifier,
		})
	}
	return mechanisms
}
