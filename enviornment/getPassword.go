package enviornment

import (
	"os"

	"github.com/programmierigel/pwmanager/logger"
)

func Password(defaultPassword string) string {
	password := os.Getenv("PASSWORD")

	if password == "" {
		logger.Info("Using default password")
		return defaultPassword
	}
	logger.Critiacal("Using env password")
	return password
}
