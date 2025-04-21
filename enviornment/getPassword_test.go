package enviornment_test

import (
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

		passwordToCheck := enviornment.Path("")

		assert.Equal(t, password, passwordToCheck)

		os.Setenv("LOCATION_PATH", oldPassword)
	})

	t.Run("Returns the default path", func(t *testing.T) {
		oldPath := os.Getenv("LOCATION_PATH")
		path := "any path"

		os.Unsetenv("LOCATION_PATH")

		pathToCheck := enviornment.Path(path)

		assert.Equal(t, path, pathToCheck)

		os.Setenv("LOCATION_PATH", oldPath)
	})
}
