package logger

func (l *Logger) Printf(format string, v ...any) {
	l.loggerFile.Printf(format, v...)
	l.loggerFile.Printf(format, v...)
}
