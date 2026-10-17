package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"email-server/smtp/internal/dkim"
	"email-server/smtp/internal/queue/internal/retry"
	"email-server/smtp/internal/queue/internal/storage"
	"email-server/smtp/internal/queue/internal/workers"
	"email-server/smtp/internal/queue/pkg/logger"

	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("SMTP_POSTGRES_DSN")
	if dsn == "" {
		logger.Error("SMTP_POSTGRES_DSN is required")
		os.Exit(1)
	}
	interval, err := durationEnv("SMTP_QUEUE_INTERVAL", 30*time.Second)
	if err != nil {
		logger.Error("queue configuration error: %v", err)
		os.Exit(1)
	}
	timeout, err := durationEnv("SMTP_DELIVERY_TIMEOUT", 10*time.Second)
	if err != nil {
		logger.Error("queue configuration error: %v", err)
		os.Exit(1)
	}
	localName := os.Getenv("SMTP_HOSTNAME")
	if localName == "" {
		localName = "localhost"
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("open postgres: %v", err)
		os.Exit(1)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := db.PingContext(ctx); err != nil {
		cancel()
		logger.Error("connect to postgres: %v", err)
		os.Exit(1)
	}
	cancel()

	queueStore := storage.NewPostgresQueueStorage(db)
	sender := workers.NewSender(localName, timeout)
	keyPath := os.Getenv("SMTP_DKIM_KEY_PATH")
	dkimDomain := os.Getenv("SMTP_DKIM_DOMAIN")
	dkimSelector := os.Getenv("SMTP_DKIM_SELECTOR")
	if keyPath != "" || dkimDomain != "" || dkimSelector != "" {
		if keyPath == "" || dkimDomain == "" || dkimSelector == "" {
			logger.Error("SMTP_DKIM_KEY_PATH, SMTP_DKIM_DOMAIN, and SMTP_DKIM_SELECTOR must be set together")
			os.Exit(1)
		}
		key, err := dkim.LoadKey(keyPath)
		if err != nil {
			logger.Error("load DKIM key: %v", err)
			os.Exit(1)
		}
		sender.Signer = &dkim.Signer{Domain: dkimDomain, Selector: dkimSelector, Key: key}
	}
	dispatcher := workers.NewDispatcher(sender, queueStore)
	retryManager := retry.NewRetryManager(queueStore)
	worker := workers.NewWorker(queueStore, dispatcher, retryManager, interval)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	logger.Info("SMTP queue worker started")
	worker.Run(ctx)
}

func durationEnv(name string, fallback time.Duration) (time.Duration, error) {
	value := os.Getenv(name)
	if value == "" {
		return fallback, nil
	}
	duration, err := time.ParseDuration(value)
	if err != nil || duration <= 0 {
		return 0, fmt.Errorf("%s must be a positive duration", name)
	}
	return duration, nil
}
