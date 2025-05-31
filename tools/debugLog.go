package tools

import (
	"fmt"
	"net/http"

	"github.com/programmierigel/pwmanager/logger"
)

func DebugLog(s string, request *http.Request, logger *logger.Logger) {
	hostPart := fmt.Sprintf("Run by Host %s (RemoteAddr: %s,\n Proto: %s,\n Pattern: %s,\n URL: %s,\n ReqURI: %s).", request.Host, request.RemoteAddr, request.Proto, request.Pattern, request.URL, request.RequestURI)
	logger.Printf("[DEBUG]-[API] %s\n%s", s, hostPart)
}
