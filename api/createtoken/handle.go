package createtoken

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
			tools.WarningLog("Attempt to create new token. Cant read request.", err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			tools.WarningLog("Attempt get url. Cant unmarshal request.", err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		token, err := store.CreateToken(requestBody.AccountName, requestBody.Password)

		if err != nil {
			tools.WarningLog(fmt.Sprintf("Attempt to create new token (%s). Cant create token.", requestBody.AccountName), err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody := ResponseBody{
			Token: token,
		}

		responseBytes, err := json.Marshal(responseBody)
		if err != nil {
			tools.WarningLog(fmt.Sprintf("Attempt to create new token(%s). Cant marshal token struct.", requestBody.AccountName), err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		tools.DebugLog(fmt.Sprintf("Successfully new token on account %s created", requestBody.AccountName), request)
		response.Write(responseBytes)
	}
}
