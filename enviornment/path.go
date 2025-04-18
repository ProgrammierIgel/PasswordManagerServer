package enviornment

import (
	"fmt"
	"os"

	"github.com/programmierigel/pwmanager/logger"
)

func Path(defaultPath string) string {
	path := os.Getenv("LOCATION_PATH")

	if path == "" {
		logger.Info(fmt.Sprintf("[Enviornment Vars] Path is set to default path: '%s'", defaultPath))
		return defaultPath
	}

	logger.Info(fmt.Sprintf("[Enviornment Vars] Path is set to env path: '%s'", path))
	return path
}
