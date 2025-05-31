package enviornment_test

import (
	"log"
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

		f, err := os.OpenFile("./test.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		passwordToCheck := enviornment.Password("", logger)

		assert.Equal(t, password, passwordToCheck)

		os.Setenv("PASSWORD", oldPassword)
		os.Remove("./test.log")
	})

	t.Run("Returns the default password", func(t *testing.T) {
		oldPassword := os.Getenv("PASSWORD")
		password := "any password"

		os.Unsetenv("PASSWORD")
		f, err := os.OpenFile("./test.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		passwordToCheck := enviornment.Password(password, logger)

		assert.Equal(t, password, passwordToCheck)

		os.Setenv("PASSWORD", oldPassword)
		os.Remove("./test.log")
	})
}
