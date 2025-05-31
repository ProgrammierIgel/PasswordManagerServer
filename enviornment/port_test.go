package enviornment_test

import (
	"log"
	"os"
	"testing"

	"fmt"

	"github.com/programmierigel/pwmanager/enviornment"
	"github.com/stretchr/testify/assert"
)

func TestPort(t *testing.T) {
	t.Run("Returns the port from enviornment variable", func(t *testing.T) {
		oldPort := os.Getenv("PORT")
		port := 1234

		os.Setenv("PORT", fmt.Sprint(port))
		f, err := os.OpenFile("./test.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		portToCheck := enviornment.Port(246, logger)

		assert.Equal(t, port, portToCheck)
		os.Setenv("PORT", oldPort)
	})

	t.Run("Returns the default port", func(t *testing.T) {
		oldPort := os.Getenv("PORT")
		port := 1234

		os.Unsetenv("PORT")
		f, err := os.OpenFile("./test.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		portToCheck := enviornment.Port(port, logger)

		assert.Equal(t, port, portToCheck)
		os.Setenv("PORT", oldPort)
	})

}
