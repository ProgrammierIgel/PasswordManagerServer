package enviornment

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func Port(defaultPort int, logger *log.Logger) int {
	portAsString := os.Getenv("PORT")

	if portAsString == "" {
		logger.Printf("[INFO]-[Enviornment Vars] Server listen on default port %d", defaultPort)
		return defaultPort
	}

	port, err := strconv.Atoi(portAsString)
	if err != nil {
		return defaultPort
	}
	fmt.Println(port)
	logger.Printf("[INFO]-[Enviornment Vars] Server listen on env port %d", port)
	return port
}
