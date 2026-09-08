package config

import "os"

type Config struct {
	HTTPAddr     string
	DatabaseURL  string
	ChatAddr     string
	MLServiceURL string
	JWTSecret    string

	// MinIO storage configuration
	MinioEndpoint  string
	MinioAccessKey string
	MinioSecretKey string
	MinioBucket    string
	MinioUseSSL    bool
}

func Load() Config {
	return Config{
		HTTPAddr:     get("HTTP_ADDR", ":8080"),
		DatabaseURL:  get("DATABASE_URL", "postgres://edu247_user:password@localhost:5432/edu247?sslmode=disable"),
		ChatAddr:     get("CHAT_ADDR", "localhost:9000"),
		MLServiceURL: get("ML_URL", "http://localhost:5000"),
		JWTSecret:    get("JWT_SECRET", "dev-secret"),

		MinioEndpoint:  get("MINIO_ENDPOINT", "localhost:9001"),
		MinioAccessKey: get("MINIO_ACCESS_KEY", "minioadmin"),
		MinioSecretKey: get("MINIO_SECRET_KEY", "minioadmin"),
		MinioBucket:    get("MINIO_BUCKET", "edu247"),
		MinioUseSSL:    get("MINIO_USE_SSL", "false") == "true",
	}
}

func get(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
