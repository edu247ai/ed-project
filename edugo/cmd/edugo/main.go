package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"edu247/edugo/internal/config"
	"edu247/edugo/internal/db"
	httpapi "edu247/edugo/internal/http"
	"edu247/edugo/internal/storage"
	"edu247/edugo/internal/workers"
)

func main() {

	// ------------------------------------------------------------
	// RUN PYTHON DB MIGRATIONS
	// ------------------------------------------------------------
	log.Println("Running DB migrations...")
	migrate := exec.Command("python3", "model/db/migrate.py")
	migrate.Stdout = os.Stdout
	migrate.Stderr = os.Stderr

	if err := migrate.Run(); err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}

	// ------------------------------------------------------------
	// LOAD CONFIG
	// ------------------------------------------------------------
	cfg := config.Load()

	log.Println("========================================")
	log.Println("        EDU247 Backend Starting")
	log.Println("========================================")

	// ------------------------------------------------------------
	// DATABASE
	// ------------------------------------------------------------
	database, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("DB init failed: %v", err)
	}
	defer database.Close()

	// ------------------------------------------------------------
	// STORAGE LAYERS
	// ------------------------------------------------------------
	local := storage.NewLocal("uploads")

	minio, err := storage.NewMinio(
		cfg.MinioEndpoint,
		cfg.MinioAccessKey,
		cfg.MinioSecretKey,
		cfg.MinioBucket,
		cfg.MinioUseSSL,
	)
	if err != nil {
		log.Fatalf("MinIO init failed: %v", err)
	}

	cstore := storage.NewCStorage()

	storeManager := &storage.Manager{
		Local: local,
		Minio: minio,
		C:     cstore,
	}

	// ------------------------------------------------------------
	// ROUTER + MIDDLEWARE PACK
	// ------------------------------------------------------------
	router := httpapi.NewRouter(database.Pool, cfg, storeManager)

	srv := &http.Server{
		Addr:         cfg.HTTPAddr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  90 * time.Second,
	}

	// ------------------------------------------------------------
	// BACKGROUND WORKERS
	// ------------------------------------------------------------
	supervisor := workers.NewSupervisor()
	supervisor.Register(workers.NewCleanupWorker())
	supervisor.Register(workers.NewAnalyticsWorker())
	supervisor.Register(workers.NewEmailWorker())

	workerCtx, workerCancel := context.WithCancel(context.Background())
	go supervisor.Start(workerCtx)

	// ------------------------------------------------------------
	// START SERVER
	// ------------------------------------------------------------
	go func() {
		log.Printf("EDU247 API listening on %s", cfg.HTTPAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// ------------------------------------------------------------
	// GRACEFUL SHUTDOWN
	// ------------------------------------------------------------
	waitForShutdown()

	log.Println("Shutting down server gracefully...")

	// Stop workers first
	workerCancel()

	// Shutdown HTTP server
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("Graceful shutdown failed: %v", err)
	} else {
		log.Println("Server stopped cleanly.")
	}

	log.Println("========================================")
	log.Println("        EDU247 Backend Stopped")
	log.Println("========================================")
}

func waitForShutdown() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
