package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// SMTPConfig holds all runtime configuration for your SMTP server.
type SMTPConfig struct {
	Hostname       string
	ListenAddress  string
	ReadTimeout    time.Duration // per command
	WriteTimeout   time.Duration
	MaxMessageSize int
	MaxConnections int
	EnableSPF      bool
	EnableDKIM     bool
	EnableDMARC    bool
	EnableSpam     bool
	QueueInterval  time.Duration
	PostgresDSN    string
	RspamdURL      string
	DKIMKeyPath    string
	DKIMDomain     string
	DKIMSelector   string
}

func Load() (SMTPConfig, error) {
	cfg := SMTPConfig{
		Hostname:       env("SMTP_HOSTNAME", "localhost"),
		ListenAddress:  env("SMTP_LISTEN_ADDRESS", ":2525"),
		ReadTimeout:    5 * time.Minute,
		WriteTimeout:   30 * time.Second,
		MaxMessageSize: 20 << 20,
		MaxConnections: 500,
		EnableSPF:      true,
		EnableDKIM:     true,
		EnableDMARC:    true,
		EnableSpam:     false,
		PostgresDSN:    os.Getenv("SMTP_POSTGRES_DSN"),
		RspamdURL:      env("SMTP_RSPAMD_URL", "http://127.0.0.1:11333/checkv2"),
		DKIMKeyPath:    os.Getenv("SMTP_DKIM_KEY_PATH"),
		DKIMDomain:     os.Getenv("SMTP_DKIM_DOMAIN"),
		DKIMSelector:   os.Getenv("SMTP_DKIM_SELECTOR"),
		QueueInterval:  30 * time.Second,
	}

	var err error
	if cfg.ReadTimeout, err = durationEnv("SMTP_READ_TIMEOUT", cfg.ReadTimeout); err != nil {
		return SMTPConfig{}, err
	}
	if cfg.WriteTimeout, err = durationEnv("SMTP_WRITE_TIMEOUT", cfg.WriteTimeout); err != nil {
		return SMTPConfig{}, err
	}
	if cfg.QueueInterval, err = durationEnv("SMTP_QUEUE_INTERVAL", cfg.QueueInterval); err != nil {
		return SMTPConfig{}, err
	}
	if cfg.MaxMessageSize, err = intEnv("SMTP_MAX_MESSAGE_SIZE", cfg.MaxMessageSize); err != nil {
		return SMTPConfig{}, err
	}
	if cfg.MaxConnections, err = intEnv("SMTP_MAX_CONNECTIONS", cfg.MaxConnections); err != nil {
		return SMTPConfig{}, err
	}
	if cfg.EnableSPF, err = boolEnv("SMTP_ENABLE_SPF", cfg.EnableSPF); err != nil {
		return SMTPConfig{}, err
	}
	if cfg.EnableDMARC, err = boolEnv("SMTP_ENABLE_DMARC", cfg.EnableDMARC); err != nil {
		return SMTPConfig{}, err
	}
	if cfg.EnableDKIM, err = boolEnv("SMTP_ENABLE_DKIM", cfg.EnableDKIM); err != nil {
		return SMTPConfig{}, err
	}
	if cfg.EnableSpam, err = boolEnv("SMTP_ENABLE_SPAM", cfg.EnableSpam); err != nil {
		return SMTPConfig{}, err
	}
	if cfg.Hostname == "" || cfg.ListenAddress == "" || cfg.ReadTimeout <= 0 ||
		cfg.WriteTimeout <= 0 || cfg.MaxMessageSize <= 0 || cfg.MaxConnections <= 0 {
		return SMTPConfig{}, fmt.Errorf("SMTP hostname, listen address, timeouts, and limits must be valid")
	}
	return cfg, nil
}

func env(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}

func intEnv(name string, fallback int) (int, error) {
	value := os.Getenv(name)
	if value == "" {
		return fallback, nil
	}
	n, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", name, err)
	}
	return n, nil
}

func durationEnv(name string, fallback time.Duration) (time.Duration, error) {
	value := os.Getenv(name)
	if value == "" {
		return fallback, nil
	}
	d, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", name, err)
	}
	return d, nil
}

func boolEnv(name string, fallback bool) (bool, error) {
	value := os.Getenv(name)
	if value == "" {
		return fallback, nil
	}
	enabled, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("%s: %w", name, err)
	}
	return enabled, nil
}
