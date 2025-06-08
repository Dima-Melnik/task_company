package utils

import (
	"log"
	"os"
	"tasks_company/app/utils/logger"

	"github.com/joho/godotenv"
)

func InitENV(secretStr string) string {
	err := godotenv.Load()
	if err == nil {
		logger.LogWithFieldsFatalf("env", "loading", err)
	}

	secret := os.Getenv(secretStr)
	if secret == "" {
		log.Fatal("Error getting secret value from env")
	}

	return secret
}
