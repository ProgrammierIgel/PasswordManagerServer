package enviornment

import (
	"os"
	"strconv"

	"github.com/programmierigel/pwmanager/logger"
)

func Port(defaultPort int, logger *logger.Logger) int {
	portAsString := os.Getenv("PORT")

	if portAsString == "" {
		logger.Printf("[INFO]-[Enviornment Vars] Server listen on default port %d", defaultPort)
		return defaultPort
	}

	port, err := strconv.Atoi(portAsString)
	if err != nil {
		return defaultPort
	}
	logger.Printf("[INFO]-[Enviornment Vars] Server listen on env port %d", port)
	return port
}
