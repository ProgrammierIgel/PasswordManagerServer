package enviornment

import (
	"os"

	"github.com/programmierigel/pwmanager/logger"
)

func Password(defaultPassword string) string {
	password := os.Getenv("PASSWORD")

	if password == "" {
		logger.Info("[Enviornment Vars] Using default password")
		return defaultPassword
	}
	logger.Critiacal("[Enviornment Vars] Using env password")
	return password
}
