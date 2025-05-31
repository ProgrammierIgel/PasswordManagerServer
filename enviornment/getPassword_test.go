package enviornment_test

import (
	"log"
	"os"
	"testing"

	"github.com/programmierigel/pwmanager/enviornment"
	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	t.Run("Returns the path from enviornment variable", func(t *testing.T) {
		oldPassword := os.Getenv("LOCATION_PATH")
		password := "any password"

		os.Setenv("LOCATION_PATH", password)

		f, err := os.OpenFile("./test.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		passwordToCheck := enviornment.Path("", logger)

		assert.Equal(t, password, passwordToCheck)

		os.Setenv("LOCATION_PATH", oldPassword)
		os.Remove("./test.log")
	})

	t.Run("Returns the default path", func(t *testing.T) {
		oldPath := os.Getenv("LOCATION_PATH")
		path := "any path"

		os.Unsetenv("LOCATION_PATH")

		f, err := os.OpenFile("./test.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		logger := log.New(f, "", log.LstdFlags)
		pathToCheck := enviornment.Path(path, logger)

		assert.Equal(t, path, pathToCheck)

		os.Setenv("LOCATION_PATH", oldPath)
		os.Remove("./test.log")
	})
}
