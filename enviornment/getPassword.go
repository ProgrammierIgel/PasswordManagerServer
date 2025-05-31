package enviornment

import (
	"log"
	"os"
)

func Password(defaultPassword string, logger *log.Logger) string {
	password := os.Getenv("PASSWORD")

	if password == "" {
		logger.Printf("[INFO]-[Enviornment Vars] Using default password")
		return defaultPassword
	}
	logger.Printf("[CRITICAL]-[Enviornment Vars] Using env password")
	return password
}
