package getnumberofallregisteredtokens

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/programmierigel/pwmanager/storage"
	"github.com/programmierigel/pwmanager/tools"
)

func Handle(store storage.Store) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		requestBytes, err := io.ReadAll(io.LimitReader(request.Body, 4096))
		if err != nil {
			tools.WarningLog("Attempt to get all active tokens of account. Cant read request.", err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			tools.WarningLog("Attempt to get all active tokens of account. Cant unmarshal request.", err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		tokensNumber, err := store.GetAllActiveTokens(requestBody.Token)

		if err != nil {
			tools.WarningLog("Attempt to get all active tokens. Cant get passwords.", err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody := ResponseBody{
			TokensNumber: tokensNumber,
		}

		responseBytes, err := json.Marshal(responseBody)
		if err != nil {
			tools.WarningLog("Attempt to get all active tokens of account. Cant marshal password struct.", err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		tools.DebugLog("Getted all active tokens of account.", request)
		response.Write(responseBytes)
	}
}
