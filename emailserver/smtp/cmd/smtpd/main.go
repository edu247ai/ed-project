package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"email-server/smtp/internal/dkim"
	"email-server/smtp/internal/dmarc"
	"email-server/smtp/internal/handlers"
	"email-server/smtp/internal/server"
	"email-server/smtp/internal/spam"
	"email-server/smtp/internal/spf"
	"email-server/smtp/internal/storage"
	"email-server/smtp/pkg/config"
	"email-server/smtp/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		logger.Error("configuration error: %v", err)
		os.Exit(1)
	}

	var store handlers.MessageStorage
	if cfg.PostgresDSN == "" {
		logger.Info("SMTP_POSTGRES_DSN is unset; using volatile in-memory storage")
		store = storage.NewMemoryStorage()
	} else {
		pgStore, err := storage.NewPostgresStorage(cfg.PostgresDSN)
		if err != nil {
			logger.Error("storage initialization failed: %v", err)
			os.Exit(1)
		}
		defer pgStore.Close()
		store = pgStore
	}

	deps := handlers.Dependencies{Storage: store}
	if cfg.EnableSPF {
		deps.SPF = spf.EvaluateContext
	}
	if cfg.EnableDKIM {
		deps.DKIM = dkim.VerifyMessage
	}
	if cfg.EnableDMARC {
		deps.DMARC = dmarc.EvaluateAlignedContext
	}
	if cfg.EnableSpam {
		deps.Spam = spam.NewRspamdClient(cfg.RspamdURL)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	srv := server.NewWithDependencies(cfg, deps)
	if err := srv.ListenAndServe(ctx); err != nil {
		logger.Error("SMTP server stopped: %v", err)
		os.Exit(1)
	}
}
