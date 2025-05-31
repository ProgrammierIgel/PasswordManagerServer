package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

func New(path string) *Logger {
	now := time.Now()
	dirPath := os.Getenv("LOCATION_PATH")

	if dirPath == "" {
		dirPath = "."
	}

	if path == "" {
		path = fmt.Sprintf("%s/PWManagerServer - %d.%d.%d.log", dirPath, now.Day(), now.Month(), now.Year())
	}

	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	loggerFile := log.New(f, "", log.LstdFlags)
	loggerConsole := log.New(os.Stdout, "", log.LstdFlags)

	return &Logger{
		loggerFile:    loggerFile,
		loggerConsole: loggerConsole,
	}
}
