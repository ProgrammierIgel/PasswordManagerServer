package logger

func (l *Logger) Println(v ...any) {
	l.loggerConsole.Println(v...)
	l.loggerFile.Println(v...)
}
