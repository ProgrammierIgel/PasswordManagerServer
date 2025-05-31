package enviornment_test

import (
	"os"
	"testing"

	"github.com/programmierigel/pwmanager/enviornment"
	"github.com/programmierigel/pwmanager/logger"
	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	t.Run("Returns the password from enviornment variable", func(t *testing.T) {
		oldPassword := os.Getenv("PASSWORD")
		password := "any password"

		os.Setenv("PASSWORD", password)

		logger := logger.New("./test.json")
		passwordToCheck := enviornment.Password("", logger)

		assert.Equal(t, password, passwordToCheck)

		os.Setenv("PASSWORD", oldPassword)
		os.Remove("./test.log")
	})

	t.Run("Returns the default password", func(t *testing.T) {
		oldPassword := os.Getenv("PASSWORD")
		password := "any password"

		os.Unsetenv("PASSWORD")
		logger := logger.New("./test.json")
		passwordToCheck := enviornment.Password(password, logger)

		assert.Equal(t, password, passwordToCheck)

		os.Setenv("PASSWORD", oldPassword)
		os.Remove("./test.log")
	})
}
