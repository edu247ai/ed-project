package secrets

import (
	"encoding/base64"
	"errors"
	"os"
)

func LoadKey(envName string) ([]byte, error) {
	raw := os.Getenv(envName)
	if raw == "" {
		return nil, errors.New("missing secret: " + envName)
	}

	key, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, err
	}

	if len(key) != 32 {
		return nil, errors.New("secret must be 32 bytes for AES-256")
	}

	return key, nil
}
