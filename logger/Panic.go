package logger

func (l *Logger) Panic(v ...any) {
	l.loggerConsole.Println(v...)
	l.loggerFile.Panic(v...)
}
