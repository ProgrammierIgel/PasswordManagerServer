package issyncdisabled

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/programmierigel/pwmanager/logger"
	"github.com/programmierigel/pwmanager/storage"
	"github.com/programmierigel/pwmanager/tools"
)

func Handle(store storage.Store) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		status := store.IsSyncDisabled()

		responseBody := ResponseBody{
			Status: status,
		}

		responseBytes, err := json.Marshal(responseBody)
		tools.WarningLog("Attempt to get is sync disabled variable. Cant marshal password struct. SYNC IS STILL ENABLED", err, request)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		hostPart := fmt.Sprintf("Run by Host %s (RemoteAddr: %s,\n Proto: %s,\n Pattern: %s,\n URL: %s,\n ReqURI: %s).", request.Host, request.RemoteAddr, request.Proto, request.Pattern, request.URL, request.RequestURI)
		logger.Info(fmt.Sprintf("GET IF SYNC IS DISABLED: %s", hostPart))
		response.Write(responseBytes)
	}
}
