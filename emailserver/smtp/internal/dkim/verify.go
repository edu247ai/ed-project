package dkim

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"net"
	"net/mail"
	"strings"
)

func VerifyMessage(ctx context.Context, raw []byte) (string, bool, error) {
	message, err := mail.ReadMessage(bytes.NewReader(raw))
	if err != nil {
		return "", false, nil
	}
	signatures := message.Header["Dkim-Signature"]
	if len(signatures) == 0 {
		return "", false, nil
	}
	body, err := readMessageBody(message)
	if err != nil {
		return "", false, err
	}
	for _, signature := range signatures {
		tags := parseTags(signature)
		domain := strings.TrimSpace(tags["d"])
		selector := strings.TrimSpace(tags["s"])
		if domain == "" || selector == "" {
			continue
		}
		key, err := lookupPublicKey(ctx, domain, selector)
		if err != nil {
			var dnsErr *net.DNSError
			if errors.As(err, &dnsErr) && dnsErr.IsNotFound {
				continue
			}
			return "", false, err
		}
		if key != nil && verifyWithKey(message.Header, body, tags, key) {
			return domain, true, nil
		}
	}
	return "", false, nil
}

func readMessageBody(message *mail.Message) ([]byte, error) {
	var body bytes.Buffer
	_, err := body.ReadFrom(message.Body)
	return body.Bytes(), err
}

func lookupPublicKey(ctx context.Context, domain, selector string) (*rsa.PublicKey, error) {
	records, err := net.DefaultResolver.LookupTXT(ctx, selector+"._domainkey."+domain)
	if err != nil {
		return nil, err
	}
	for _, record := range records {
		tags := parseTags(record)
		if strings.EqualFold(tags["k"], "rsa") || tags["k"] == "" {
			encoded := strings.TrimSpace(tags["p"])
			if encoded == "" {
				continue
			}
			der, err := base64.StdEncoding.DecodeString(encoded)
			if err != nil {
				continue
			}
			publicKey, err := x509.ParsePKIXPublicKey(der)
			if err == nil {
				if rsaKey, ok := publicKey.(*rsa.PublicKey); ok {
					return rsaKey, nil
				}
			}
			if rsaKey, err := x509.ParsePKCS1PublicKey(der); err == nil {
				return rsaKey, nil
			}
		}
	}
	return nil, nil
}

func verifyWithKey(headers mail.Header, body []byte, tags map[string]string, key *rsa.PublicKey) bool {
	if !strings.EqualFold(tags["a"], "rsa-sha256") || !strings.EqualFold(tags["c"], "relaxed/simple") {
		return false
	}
	bodyHash, err := base64.StdEncoding.DecodeString(tags["bh"])
	if err != nil {
		return false
	}
	actualBodyHash := sha256.Sum256([]byte(CanonicalizeBody(body)))
	if !bytes.Equal(bodyHash, actualBodyHash[:]) {
		return false
	}
	signature, err := base64.StdEncoding.DecodeString(tags["b"])
	if err != nil {
		return false
	}
	var signingData strings.Builder
	for _, name := range strings.Split(tags["h"], ":") {
		value := headers.Get(strings.TrimSpace(name))
		if value == "" {
			return false
		}
		signingData.WriteString(canonicalizeHeader(name, value))
	}
	signingData.WriteString(canonicalizeHeader("DKIM-Signature", withoutSignatureValue(tags["raw"])))
	hash := sha256.Sum256([]byte(signingData.String()))
	return rsa.VerifyPKCS1v15(key, crypto.SHA256, hash[:], signature) == nil
}

func withoutSignatureValue(header string) string {
	parts := strings.Split(header, ";")
	for i, part := range parts {
		name, _, ok := strings.Cut(strings.TrimSpace(part), "=")
		if ok && strings.EqualFold(name, "b") {
			parts[i] = " b="
		}
	}
	return strings.Join(parts, ";")
}

func parseTags(value string) map[string]string {
	tags := make(map[string]string)
	for _, part := range strings.Split(value, ";") {
		name, content, ok := strings.Cut(part, "=")
		if ok {
			tags[strings.ToLower(strings.TrimSpace(name))] = strings.TrimSpace(content)
		}
	}
	tags["raw"] = value
	return tags
}
