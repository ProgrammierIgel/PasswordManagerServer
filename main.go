package main

import (
	"fmt"
	"net/http"

	"github.com/programmierigel/pwmanager/api"
	"github.com/programmierigel/pwmanager/enviornment"
	"github.com/programmierigel/pwmanager/logger"
	"github.com/programmierigel/pwmanager/storage/inmemory"
)

func main() {

	logger := logger.New("")
	port := enviornment.Port(3000, logger)

	path := enviornment.Path(".", logger)
	password := enviornment.Password("123", logger)

	store := inmemory.New(path, password, logger)

	router := api.GetRouter(store, logger)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Panic(err)
	}
}
