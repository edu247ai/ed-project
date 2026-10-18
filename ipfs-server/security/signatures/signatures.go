package signatures

import (
	"crypto/hmac"
	"crypto/sha256"
)

type Signer struct {
	key []byte
}

func New(key []byte) *Signer {
	return &Signer{key: key}
}

func (s *Signer) Sign(data []byte) []byte {
	mac := hmac.New(sha256.New, s.key)
	mac.Write(data)
	return mac.Sum(nil)
}

func (s *Signer) Verify(data, signature []byte) bool {
	mac := hmac.New(sha256.New, s.key)
	mac.Write(data)
	expected := mac.Sum(nil)
	return hmac.Equal(expected, signature)
}
