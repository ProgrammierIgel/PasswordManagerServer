package getallpasswordsofaccount

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/programmierigel/pwmanager/logger"
	"github.com/programmierigel/pwmanager/storage"
	"github.com/programmierigel/pwmanager/tools"
)

func Handle(store storage.Store, logger *logger.Logger) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		requestBytes, err := io.ReadAll(io.LimitReader(request.Body, 4096))
		if err != nil {
			tools.WarningLog("Attempt to get all passwordnames of account. Cant read request.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			tools.WarningLog("Attempt to get all passwordnames of account. Cant unmarshal request.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		passwordNames, err := store.GetAllPasswordNamesOfAccount(requestBody.Token)

		if err != nil {
			tools.WarningLog("Attempt to get all passwordnames. Cant get passwords.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody := ResponseBody{
			PasswordNames: passwordNames,
		}

		responseBytes, err := json.Marshal(responseBody)
		if err != nil {
			tools.WarningLog("Attempt to get all passwordnames of account. Cant marshal password struct.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		tools.DebugLog("Getted all passwordnames of account.", request, logger)
		response.Write(responseBytes)
	}
}
