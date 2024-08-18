package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    DBSSLMode  string
)

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    DBHost = getEnv("DB_HOST", "localhost")
    DBPort = getEnv("DB_PORT", "5432")
    DBUser = getEnv("DB_USER", "postgres")
    DBPassword = getEnv("DB_PASSWORD", "password")
    DBName = getEnv("DB_NAME", "jobtracker")
    DBSSLMode = getEnv("DB_SSLMODE", "disable")
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
