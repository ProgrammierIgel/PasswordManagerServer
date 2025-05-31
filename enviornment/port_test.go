package enviornment_test

import (
	"os"
	"testing"

	"fmt"

	"github.com/programmierigel/pwmanager/enviornment"
	"github.com/programmierigel/pwmanager/logger"
	"github.com/stretchr/testify/assert"
)

func TestPort(t *testing.T) {
	t.Run("Returns the port from enviornment variable", func(t *testing.T) {
		oldPort := os.Getenv("PORT")
		port := 1234

		os.Setenv("PORT", fmt.Sprint(port))
		logger := logger.New("./test.log")
		portToCheck := enviornment.Port(246, logger)

		assert.Equal(t, port, portToCheck)
		os.Setenv("PORT", oldPort)
	})

	t.Run("Returns the default port", func(t *testing.T) {
		oldPort := os.Getenv("PORT")
		port := 1234

		os.Unsetenv("PORT")

		logger := logger.New("./test.json")
		portToCheck := enviornment.Port(port, logger)

		assert.Equal(t, port, portToCheck)
		os.Setenv("PORT", oldPort)
		os.Remove("./test.json")
	})

}
