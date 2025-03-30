package logger

import (
	"fmt"
	"time"

	terminalcolor "github.com/programmierigel/pwmanager/logger/terminalColor"
)

func Info(s string) {
	if GetLogLevel() == "critical" {
		return
	}
	currentTime := time.Now()
	msg := fmt.Sprintf("[INFO]: %s: "+s, currentTime.Format("15:04:05 - 02.01.2006"))

	fmt.Println(terminalcolor.SetColor(msg, "Green"))
}
