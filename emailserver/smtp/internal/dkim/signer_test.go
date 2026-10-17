package dkim

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"io"
	"net/mail"
	"strings"
	"testing"
)

func TestSignerAddsDKIMSignatureToMessage(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatal(err)
	}
	signer := &Signer{Domain: "example.test", Selector: "mail", Key: key}
	raw := []byte("From: sender@example.test\r\nSubject: test\r\nDate: Sat, 17 Oct 2026 12:00:00 +0000\r\n\r\nbody\r\n")

	signed, err := signer.SignMessage(raw)
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.SplitN(string(signed), "\r\n", 2)
	if len(lines) != 2 || !strings.HasPrefix(lines[0], "DKIM-Signature: ") {
		t.Fatalf("signed message does not start with a DKIM header: %q", signed)
	}
	if strings.Count(lines[0], "b=") != 1 {
		t.Fatalf("DKIM header should contain one signature field: %q", lines[0])
	}
	signature := strings.TrimPrefix(lines[0][strings.LastIndex(lines[0], "b=")+2:], "")
	if _, err := base64.StdEncoding.DecodeString(signature); err != nil {
		t.Fatalf("DKIM signature is not base64: %v", err)
	}
	if !strings.HasSuffix(string(signed), string(raw)) {
		t.Fatal("original message was not preserved after the DKIM header")
	}
	message, err := mail.ReadMessage(bytes.NewReader(signed))
	if err != nil {
		t.Fatal(err)
	}
	body, err := io.ReadAll(message.Body)
	if err != nil {
		t.Fatal(err)
	}
	tags := parseTags(message.Header.Get("Dkim-Signature"))
	if !verifyWithKey(message.Header, body, tags, &key.PublicKey) {
		t.Fatal("generated DKIM signature did not verify")
	}
	if verifyWithKey(message.Header, append(body, []byte("changed")...), tags, &key.PublicKey) {
		t.Fatal("modified message body unexpectedly verified")
	}
}

func TestCanonicalizeBodyUsesCRLF(t *testing.T) {
	if got := CanonicalizeBody([]byte("one\r\ntwo\r\n\r\n")); got != "one\r\ntwo\r\n" {
		t.Fatalf("CanonicalizeBody() = %q", got)
	}
	if got := CanonicalizeBody(nil); got != "\r\n" {
		t.Fatalf("CanonicalizeBody(empty) = %q", got)
	}
}
