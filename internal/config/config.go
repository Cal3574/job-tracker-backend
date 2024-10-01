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
	// Check if we're running in production by checking an "ENV" variable
	if os.Getenv("ENV") != "production" {
		// Load .env file only if we're not in production
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, using system environment variables")
		}
	} else {
		log.Println("Running in production, not loading .env file")
	}

	// Load environment variables (from .env or system environment)
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
