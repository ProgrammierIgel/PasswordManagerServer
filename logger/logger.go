package logger

import "log"

type Logger struct {
	loggerFile    *log.Logger
	loggerConsole *log.Logger
}
