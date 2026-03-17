package config

import "os"

type Config struct {
	Env          string
	Port         string
	PostgresHost string
	PostgresPort string
	PostgresDB   string
	PostgresUser string
	PostgresPass string
	OCRBaseURL   string
}

func Load() Config {
	return Config{
		Env:          getEnv("ENV", "development"),
		Port:         getEnv("GO_API_PORT", "8080"),
		PostgresHost: getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort: getEnv("POSTGRES_PORT", "5432"),
		PostgresDB:   getEnv("POSTGRES_DB", "struck_ocr"),
		PostgresUser: getEnv("POSTGRES_USER", "struck"),
		PostgresPass: getEnv("POSTGRES_PASSWORD", "struckpass"),
		OCRBaseURL:   getEnv("OCR_SERVICE_BASE_URL", "http://localhost:8000"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
