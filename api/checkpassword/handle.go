package checkpassword

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
			tools.WarningLog("Attempt to check password. Cant read request.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			tools.WarningLog("Attempt check password. Cant unmarshal request.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		err = store.CheckPassword(requestBody.AccountName, requestBody.MasterPassword, request.RemoteAddr)
		var responseBody ResponseBody
		if err != nil {
			tools.WarningLog(fmt.Sprintf("Attempt check password on account %s. Password is wrong. (%s)", requestBody.AccountName, requestBody.MasterPassword), err, request, logger)
			responseBody = ResponseBody{
				Status: false,
			}
		} else {
			responseBody = ResponseBody{
				Status: true,
			}
		}

		responseBytes, err := json.Marshal(responseBody)
		if err != nil {
			tools.WarningLog(fmt.Sprintf("Attempt check password on account %s. Cant marshal password struct. Password is correct", requestBody.AccountName), err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		if responseBody.Status {
			tools.DebugLog(fmt.Sprintf("Checked password from account %s. Password is correct", requestBody.AccountName), request, logger)
		}
		response.Write(responseBytes)
	}
}
