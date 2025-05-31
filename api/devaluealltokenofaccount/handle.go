package devaluealltokenofaccount

import (
	"encoding/json"
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
			tools.WarningLog("Attempt to devalue all tokens from account. Cant read request.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			tools.WarningLog("Attempt to devalue all tokens from account. Cant unmarshal request.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		err = store.DevalueAllTokensOfAccount(requestBody.Token)

		if err != nil {
			tools.WarningLog("Attempt to devalue all tokens from account. Cant devalue token.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "text/plain")
		response.WriteHeader(http.StatusOK)
		response.Header().Set("ok", "true")
		tools.DebugLog("Successfully all tokens of account devalued.", request, logger)
		response.Write([]byte(http.StatusText(http.StatusOK)))
	}
}
