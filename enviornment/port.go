package enviornment

import (
	"fmt"
	"os"
	"strconv"

	"github.com/programmierigel/pwmanager/logger"
)

func Port(defaultPort int) (int, error) {
	portAsString := os.Getenv("PORT")

	if portAsString == "" {
		logger.Info(fmt.Sprintf("[Enviornment Vars] Server listen on default port %d", defaultPort))
		return defaultPort, nil
	}

	port, err := strconv.Atoi(portAsString)
	if err != nil {
		return 0, err
	}
	fmt.Println(port)
	logger.Info(fmt.Sprintf("[Enviornment Vars] Server listen on env port %d", port))
	return port, nil
}
