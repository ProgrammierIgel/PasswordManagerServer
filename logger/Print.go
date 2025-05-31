package logger

func (l *Logger) Print(v ...any) {
	l.loggerFile.Print(v...)
	l.loggerConsole.Print(v...)
}
