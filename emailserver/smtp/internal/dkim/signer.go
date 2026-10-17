package dkim

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"net/mail"
	"strings"
)

type Signer struct {
	Domain   string
	Selector string
	Key      *rsa.PrivateKey
}

func (s *Signer) Sign(body []byte, headers map[string]string) (string, error) {
	if s.Key == nil || s.Domain == "" || s.Selector == "" ||
		strings.ContainsAny(s.Domain+s.Selector, "\r\n; \t") {
		return "", errors.New("DKIM signer requires a key, domain, and selector")
	}
	canonBody := CanonicalizeBody(body)
	bodyHash := sha256.Sum256([]byte(canonBody))
	dkimHeader := BuildDKIMHeader(s.Domain, s.Selector, headers, bodyHash[:])
	var signedHeaders strings.Builder
	for _, name := range []string{"from", "subject", "date"} {
		if value, ok := headerValue(headers, name); ok && value != "" {
			signedHeaders.WriteString(canonicalizeHeader(name, value))
		}
	}
	signingData := signedHeaders.String() + canonicalizeHeader("DKIM-Signature", strings.TrimPrefix(dkimHeader, "DKIM-Signature:"))
	sigHash := sha256.Sum256([]byte(signingData))
	sig, err := rsa.SignPKCS1v15(rand.Reader, s.Key, crypto.SHA256, sigHash[:])
	if err != nil {
		return "", err
	}
	return dkimHeader + base64.StdEncoding.EncodeToString(sig) + "\r\n", nil
}

func (s *Signer) SignMessage(raw []byte) ([]byte, error) {
	message, err := mail.ReadMessage(bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(message.Body)
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		"from":    message.Header.Get("From"),
		"subject": message.Header.Get("Subject"),
		"date":    message.Header.Get("Date"),
	}
	if headers["from"] == "" {
		return nil, errors.New("outbound message has no From header")
	}
	signature, err := s.Sign(body, headers)
	if err != nil {
		return nil, err
	}
	result := make([]byte, 0, len(signature)+len(raw))
	result = append(result, signature...)
	result = append(result, raw...)
	return result, nil
}
