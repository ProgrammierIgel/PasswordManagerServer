package enablesync

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/programmierigel/pwmanager/storage"
	"github.com/programmierigel/pwmanager/tools"
)

func Handle(store storage.Store, logger *log.Logger) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		// TODO: Get Admin Password from env variable and check if it is correct. Than dis/enable sync and run sync to/from file. For now the code is a comment

		requestBytes, err := io.ReadAll(io.LimitReader(request.Body, 4096))
		if err != nil {
			tools.WarningLog("Attempt to enable syncronization. Cant read request.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			tools.WarningLog("Attempt to enable syncronization. Cant unmarshal request.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		status, err := store.EnableSync(requestBody.Password)
		if err != nil {
			tools.WarningLog("Attempt to enable syncronization. Wrong password", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody := ResponseBody{
			Status: status,
		}

		responseBytes, err := json.Marshal(responseBody)
		if err != nil {
			tools.WarningLog("Attempt to enable sync. Cant marshal password struct. SYNC IS STILL ENABLED", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		hostPart := fmt.Sprintf("Run by Host %s (RemoteAddr: %s,\n Proto: %s,\n Pattern: %s,\n URL: %s,\n ReqURI: %s).", request.Host, request.RemoteAddr, request.Proto, request.Pattern, request.URL, request.RequestURI)
		logger.Printf("[INFO] - [API] ENABLED SYNC: %s", hostPart)
		response.Write(responseBytes)
	}
}
