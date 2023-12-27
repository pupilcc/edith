package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	value := os.Getenv(key)
	if len(value) == 0 {
		return ""
	}
	return value
}
