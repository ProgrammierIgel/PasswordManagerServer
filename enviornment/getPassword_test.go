package enviornment_test

import (
	"os"
	"testing"

	"github.com/programmierigel/pwmanager/enviornment"
	"github.com/programmierigel/pwmanager/logger"
	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	t.Run("Returns the path from enviornment variable", func(t *testing.T) {
		oldPassword := os.Getenv("LOCATION_PATH")
		password := "any password"

		os.Setenv("LOCATION_PATH", password)

		logger := logger.New("./test.json")
		passwordToCheck := enviornment.Path("", logger)

		assert.Equal(t, password, passwordToCheck)

		os.Setenv("LOCATION_PATH", oldPassword)
		os.Remove("./test.log")
	})

	t.Run("Returns the default path", func(t *testing.T) {
		oldPath := os.Getenv("LOCATION_PATH")
		path := "any path"

		os.Unsetenv("LOCATION_PATH")
		logger := logger.New("./test.json")
		pathToCheck := enviornment.Path(path, logger)

		assert.Equal(t, path, pathToCheck)

		os.Setenv("LOCATION_PATH", oldPath)
		os.Remove("./test.log")
	})
}
