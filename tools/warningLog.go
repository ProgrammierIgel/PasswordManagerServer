package tools

import (
	"fmt"
	"net/http"

	"github.com/programmierigel/pwmanager/logger"
)

func WarningLog(s string, errorMsg error, request *http.Request) {
	errorPart := fmt.Sprintf("Error: %s", errorMsg.Error())
	hostPart := fmt.Sprintf("Attemped by Host %s (RemoteAddr: %s,\n Proto: %s,\n Pattern: %s,\n URL: %s,\n ReqURI: %s).", request.Host, request.RemoteAddr, request.Proto, request.Pattern, request.URL, request.RequestURI)
	logger.Warning(fmt.Sprintf("[API] %s\n%s\n%s ", s, errorPart, hostPart))
}
