package addnewpassword

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
			tools.WarningLog("Attempt to add new password. Cant read request.", err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			tools.WarningLog("Attempt add new password. Cant unmarshal request.", err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		err = store.AddNewPassword(requestBody.MasterPassword, requestBody.AccountName, requestBody.PasswordName, requestBody.Password.Password, requestBody.Password.URL, requestBody.Password.Username)

		if err != nil {
			tools.WarningLog(fmt.Sprintf("Attempt to add new password (%s) on account %s. Cant add password.", requestBody.PasswordName, requestBody.AccountName), err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		tools.DebugLog(fmt.Sprintf("Added new password (%s) to account %s", requestBody.PasswordName, requestBody.AccountName), request)
		response.Write([]byte(http.StatusText(http.StatusOK)))
	}
}
