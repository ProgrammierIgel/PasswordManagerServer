package enviornment_test

import (
	"os"
	"testing"

	"github.com/programmierigel/pwmanager/enviornment"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestPort(t *testing.T) {
	t.Run("Returns the port from enviornment variable", func(t *testing.T) {
		oldPort := os.Getenv("PORT")
		port := 1234

		os.Setenv("PORT", fmt.Sprint(port))

		portToCheck := enviornment.Port(246)

		assert.Equal(t, port, portToCheck)
		os.Setenv("PORT", oldPort)
	})

	t.Run("Returns the default port", func(t *testing.T) {
		oldPort := os.Getenv("PORT")
		port := 1234

		os.Unsetenv("PORT")

		portToCheck := enviornment.Port(port)

		assert.Equal(t, port, portToCheck)
		os.Setenv("PORT", oldPort)
	})

}
