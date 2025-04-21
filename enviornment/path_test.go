package enviornment_test

import (
	"os"
	"testing"

	"github.com/programmierigel/pwmanager/enviornment"
	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	t.Run("Returns the password from enviornment variable", func(t *testing.T) {
		oldPassword := os.Getenv("PASSWORD")
		password := "any password"

		os.Setenv("PASSWORD", password)

		passwordToCheck := enviornment.Password("")

		assert.Equal(t, password, passwordToCheck)

		os.Setenv("PASSWORD", oldPassword)
	})

	t.Run("Returns the default password", func(t *testing.T) {
		oldPassword := os.Getenv("PASSWORD")
		password := "any password"

		os.Unsetenv("PASSWORD")

		passwordToCheck := enviornment.Password(password)

		assert.Equal(t, password, passwordToCheck)

		os.Setenv("PASSWORD", oldPassword)
	})
}
