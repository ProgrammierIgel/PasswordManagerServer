package syncfromfile

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/programmierigel/pwmanager/logger"
	"github.com/programmierigel/pwmanager/storage"
	"github.com/programmierigel/pwmanager/tools"
)

func Handle(store storage.Store, logger *logger.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := store.SyncFromFile()

		if err != nil {
			tools.WarningLog("Attempt sync cached store from file.", err, r, logger)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		status := http.StatusOK

		w.WriteHeader(status)
		w.Write([]byte(http.StatusText(status)))
	}
}
