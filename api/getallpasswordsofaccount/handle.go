package getallpasswordsofaccount

import (
	"encoding/json"
	"fmt"
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
			tools.WarningLog("Attempt to get all passwordnames of account. Cant read request.", err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			tools.WarningLog("Attempt to get all passwordnames of account. Cant unmarshal request.", err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		passwordNames, err := store.GetAllPasswordNamesOfAccount(requestBody.AccountName, requestBody.MasterPassword)

		if err != nil {
			tools.WarningLog(fmt.Sprintf("Attempt to get all passwordnames of account %s. Cant get passwords.", requestBody.AccountName), err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody := ResponseBody{
			PasswordNames: passwordNames,
		}

		responseBytes, err := json.Marshal(responseBody)
		tools.WarningLog(fmt.Sprintf("Attempt to get all passwordnames of account %s. Cant marshal password struct.", requestBody.AccountName), err, request)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		tools.DebugLog(fmt.Sprintf("Getted all passwordnames of account %s.",requestBody.AccountName), request)
		response.Write(responseBytes)
	}
}
