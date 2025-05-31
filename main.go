package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/programmierigel/pwmanager/api"
	"github.com/programmierigel/pwmanager/enviornment"
	"github.com/programmierigel/pwmanager/storage/inmemory"
)

func main() {
	now := time.Now()
	f, err := os.OpenFile(fmt.Sprintf("PWManagerServer - %d.%d.%d.logs", now.Day(), now.Month(), now.Year()),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)

	port := enviornment.Port(3000, logger)

	path := enviornment.Path(".", logger)
	password := enviornment.Password("123", logger)

	store := inmemory.New(path, password, logger)

	router := api.GetRouter(store, logger)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		logger.Panic(err)
	}
}
