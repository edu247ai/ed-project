package dkim

import "strings"

func CanonicalizeBody(body []byte) string {
	value := strings.ReplaceAll(string(body), "\r\n", "\n")
	value = strings.ReplaceAll(value, "\r", "\n")
	value = strings.TrimRight(value, "\n")
	value = strings.ReplaceAll(value, "\n", "\r\n")
	return value + "\r\n"
}
