package deletepassword

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
			tools.WarningLog("Attempt to delete password. Cant read request.", err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			tools.WarningLog("Attempt delete password. Cant unmarshal request.", err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		err = store.DeletePassword(requestBody.Token, requestBody.PasswordName)

		if err != nil {
			tools.WarningLog(fmt.Sprintf("Attempt to delete password (%s). Cant add password.", requestBody.PasswordName), err, request)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		tools.DebugLog(fmt.Sprintf("Deleted password (%s)", requestBody.PasswordName), request)
		response.Write([]byte(http.StatusText(http.StatusOK)))
	}
}
