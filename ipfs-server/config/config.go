package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	accesscontrol "ipfs-server/security/access-control"
	"ipfs-server/config/secrets"
)

type Config struct {
	API struct {
		Port           int    `yaml:"port"`
		MaxUploadBytes int64  `yaml:"max_upload_bytes"`
		Keys           []APIKey `yaml:"-"`
	} `yaml:"api"`

	Dashboard struct {
		Port int `yaml:"port"`
	} `yaml:"dashboard"`

	Swarm struct {
		Port          int      `yaml:"port"`
		BootstrapPeers []string `yaml:"bootstrap_peers"`
	} `yaml:"swarm"`

	Storage struct {
		Base string `yaml:"base"`
	} `yaml:"storage"`

	Security struct {
		EncryptionKey []byte `yaml:"-"`
		SignatureKey  []byte `yaml:"-"`
	} `yaml:"security"`
}

type APIKey struct {
	Key  string             `json:"key"`
	Role accesscontrol.Role `json:"role"`
}

func Load(profilePath string) (*Config, error) {
	data, err := os.ReadFile(profilePath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	encKey, err := secrets.LoadKey("IPFS_ENC_KEY")
	if err != nil {
		return nil, err
	}

	sigKey, err := secrets.LoadKey("IPFS_SIG_KEY")
	if err != nil {
		return nil, err
	}

	cfg.Security.EncryptionKey = encKey
	cfg.Security.SignatureKey = sigKey
	cfg.API.Keys, err = loadAPIKeys(os.Getenv("IPFS_API_KEYS"))
	if err != nil {
		return nil, err
	}
	if len(cfg.API.Keys) == 0 {
		return nil, errors.New("missing secret: IPFS_API_KEYS")
	}
	if cfg.API.MaxUploadBytes <= 0 {
		cfg.API.MaxUploadBytes = 64 << 20
	}
	if cfg.API.Port < 1 || cfg.API.Port > 65535 {
		return nil, fmt.Errorf("invalid API port %d", cfg.API.Port)
	}
	if cfg.Swarm.Port < 1 || cfg.Swarm.Port > 65535 {
		return nil, fmt.Errorf("invalid swarm port %d", cfg.Swarm.Port)
	}
	if cfg.Storage.Base == "" {
		return nil, errors.New("storage base path is required")
	}

	return &cfg, nil
}

func loadAPIKeys(raw string) ([]APIKey, error) {
	if raw == "" {
		return nil, nil
	}
	var keys []APIKey
	if err := json.Unmarshal([]byte(raw), &keys); err != nil {
		return nil, fmt.Errorf("parse IPFS_API_KEYS: %w", err)
	}
	acl := accesscontrol.New()
	for _, key := range keys {
		if err := acl.RegisterKey(key.Key, key.Role); err != nil {
			return nil, err
		}
	}
	return keys, nil
}
