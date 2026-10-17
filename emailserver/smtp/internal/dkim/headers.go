package dkim

import (
	"encoding/base64"
	"strings"
)

func BuildDKIMHeader(domain, selector string, headers map[string]string, bodyHash []byte) string {
	signedHeaders := []string{"from", "subject", "date"}
	present := make([]string, 0, len(signedHeaders))
	for _, name := range signedHeaders {
		if value, ok := headerValue(headers, name); ok && value != "" {
			present = append(present, name)
		}
	}
	return "DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/simple; d=" + domain +
		"; s=" + selector + "; h=" + strings.Join(present, ":") +
		"; bh=" + base64.StdEncoding.EncodeToString(bodyHash) + "; b="
}

func headerValue(headers map[string]string, name string) (string, bool) {
	for key, value := range headers {
		if strings.EqualFold(key, name) {
			return value, true
		}
	}
	return "", false
}

func canonicalizeHeader(name, value string) string {
	fields := strings.Fields(value)
	return strings.ToLower(name) + ":" + strings.Join(fields, " ") + "\r\n"
}
