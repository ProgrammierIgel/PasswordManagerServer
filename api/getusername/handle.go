package getusername

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
		requestBytes, err := io.ReadAll(io.LimitReader(request.Body, 4096))
		if err != nil {
			tools.WarningLog("Attempt to get a username. Cant read request.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			tools.WarningLog("Attempt get username. Cant unmarshal request.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		username, err := store.GetUsername(requestBody.Token, requestBody.PasswordName)

		if err != nil {
			tools.WarningLog(fmt.Sprintf("Attempt to get the username (%s). Cant get the username.", requestBody.PasswordName), err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody := ResponseBody{
			Username: username,
		}

		responseBytes, err := json.Marshal(responseBody)
		if err != nil {
			tools.WarningLog(fmt.Sprintf("Attempt to get the username (%s). Cant marshal username struct.", requestBody.PasswordName), err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		tools.DebugLog(fmt.Sprintf("Getted username (%s)", requestBody.PasswordName), request, logger)
		response.Write(responseBytes)
	}
}
