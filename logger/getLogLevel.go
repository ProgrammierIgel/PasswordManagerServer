package logger

import (
	"os"
)

func GetLogLevel() string {
	var DefaultLogLevel = "debug"
	var PossibleLogLevels = []string{
		"debug",
		"normal",
		"critical",
	}
	env := os.Getenv("LOG_LEVEL")

	if isElementInSlice(env, PossibleLogLevels) != nil {
		return DefaultLogLevel
	}

	return env
}
