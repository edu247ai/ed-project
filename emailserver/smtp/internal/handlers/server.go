package handlers

import (
	"context"
	"time"
)

type ServerConfig struct {
	Addr            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	Hostname        string
	MaxMessageSize  int
	EnableSPF       bool
	EnableDKIM      bool
	EnableDMARC     bool
	EnableSpam      bool
	SpamRejectScore float64
}

type MessageStorage interface {
	SaveIncomingMessage(context.Context, string, string, []string, []byte) (string, error)
}

type SpamChecker interface {
	Check(context.Context, []byte) (float64, []string, error)
}

type Dependencies struct {
	Storage MessageStorage
	Spam    SpamChecker
	SPF     func(context.Context, string, string) (string, error)
	DKIM    func(context.Context, []byte) (string, bool, error)
	DMARC   func(context.Context, string, string, string, string, string) (string, error)
}
