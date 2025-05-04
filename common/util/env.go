package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file
func LoadEnv(filePath string) {
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatalf("Error loading .env file at %s: %v", filePath, err)
	}
}

// GetEnv returns a value from environment with optional fallback
func GetEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
