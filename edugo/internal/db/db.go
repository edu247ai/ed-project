package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func Connect(databaseURL string) (*Database, error) {
	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	//
	// ------------------------------------------------------------
	// Production tuning (FIXED: cast to int32)
	// ------------------------------------------------------------
	//
	cfg.MaxConns = int32(getEnvInt("DB_MAX_CONNS", 30))
	cfg.MinConns = int32(getEnvInt("DB_MIN_CONNS", 5))
	cfg.MaxConnLifetime = 30 * time.Minute
	cfg.MaxConnIdleTime = 10 * time.Minute
	cfg.HealthCheckPeriod = 30 * time.Second

	//
	// ------------------------------------------------------------
	// Retry logic with exponential backoff
	// ------------------------------------------------------------
	//
	var pool *pgxpool.Pool
	for i := 0; i < 5; i++ {
		pool, err = pgxpool.NewWithConfig(context.Background(), cfg)
		if err == nil {
			break
		}

		wait := time.Duration(i+1) * time.Second
		log.Printf("DB connection failed (attempt %d): %v — retrying in %s", i+1, err, wait)
		time.Sleep(wait)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create database pool after retries: %w", err)
	}

	//
	// ------------------------------------------------------------
	// Validate connection
	// ------------------------------------------------------------
	//
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	log.Println("Connected to PostgreSQL successfully.")
	return &Database{Pool: pool}, nil
}

func (db *Database) Close() {
	if db.Pool != nil {
		db.Pool.Close()
		log.Println("Database connection closed.")
	}
}

// ------------------------------------------------------------
// Health check (for /api/health)
// ------------------------------------------------------------
func (db *Database) Healthy() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := db.Pool.Ping(ctx); err != nil {
		log.Printf("DB health check failed: %v", err)
		return false
	}
	return true
}

// ------------------------------------------------------------
// Helpers
// ------------------------------------------------------------
func getEnvInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	var out int
	fmt.Sscanf(val, "%d", &out)
	if out <= 0 {
		return fallback
	}
	return out
}
