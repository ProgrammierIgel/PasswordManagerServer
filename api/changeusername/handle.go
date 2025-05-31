package changeusername

import (
	"encoding/json"
	"fmt"
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
			tools.WarningLog("Attempt to change URL. Cant read request.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody RequestBody
		err = json.Unmarshal(requestBytes, &requestBody)
		if err != nil {
			tools.WarningLog("Attempt change URL. Cant unmarshal request.", err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		err = store.ChangeUsername(requestBody.Token, requestBody.PasswordName, requestBody.NewUsername)

		if err != nil {
			tools.WarningLog(fmt.Sprintf("Attempt to change URL (%s). Cant add password.", requestBody.PasswordName), err, request, logger)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		tools.DebugLog(fmt.Sprintf("Changed URL (%s)", requestBody.PasswordName), request, logger)
		response.Write([]byte(http.StatusText(http.StatusOK)))
	}
}
