package logger

import (
	"fmt"
	"time"

	terminalcolor "github.com/programmierigel/pwmanager/logger/terminalColor"
)

func Debug(s string) {
	if GetLogLevel() != "debug" {
		return
	}
	currentTime := time.Now()
	msg := fmt.Sprintf("[DEBUG]: %s: ***%s***", currentTime.Format("15:04:05 - 02.01.2006"), s)

	fmt.Println(terminalcolor.SetColor(msg, "Gray"))
}
