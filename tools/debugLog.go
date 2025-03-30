package tools

import (
	"fmt"
	"net/http"

	"github.com/programmierigel/pwmanager/logger"
)

func DebugLog(s string, request *http.Request) {
	hostPart := fmt.Sprintf("Run by Host %s (RemoteAddr: %s,\n Proto: %s,\n Pattern: %s,\n URL: %s,\n ReqURI: %s).", request.Host, request.RemoteAddr, request.Proto, request.Pattern, request.URL, request.RequestURI)
	logger.Debug(fmt.Sprintf("%s\n%s", s, hostPart))
}
