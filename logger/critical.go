package logger

import (
	"fmt"
	"time"

	terminalcolor "github.com/programmierigel/pwmanager/logger/terminalColor"
)

func Critiacal(s string) {
	currentTime := time.Now()
	msg := fmt.Sprintf("[CRITICAL]: %s: "+s, currentTime.Format("15:04:05 - 02.01.2006"))

	fmt.Println(terminalcolor.SetColor(msg, "Red"))
}
