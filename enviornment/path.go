package enviornment

import (
	"os"

	"github.com/programmierigel/pwmanager/logger"
)

func Path(defaultPath string, logger *logger.Logger) string {
	path := os.Getenv("LOCATION_PATH")

	if path == "" {
		logger.Printf("[INFO]-[Enviornment Vars] Path is set to default path: '%s'", defaultPath)
		return defaultPath
	}

	logger.Printf("[INFO]-[Enviornment Vars] Path is set to env path: '%s'", path)
	return path
}
