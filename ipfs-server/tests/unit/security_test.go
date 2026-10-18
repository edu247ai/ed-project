package unit

import (
	"ipfs-server/security"
	"testing"
)

func TestEncryptionDecryption(t *testing.T) {
	key := []byte("12345678901234567890123456789012")
	sec, _ := security.New(key, key)

	encrypted, nonce, _ := sec.Encryptor.Encrypt([]byte("hello"))
	decrypted, _ := sec.Encryptor.Decrypt(encrypted, nonce)

	if string(decrypted) != "hello" {
		t.Fatalf("expected hello, got %s", decrypted)
	}
}
