package main

import (
	"fmt"
	"net/http"

	"github.com/programmierigel/pwmanager/api"
	"github.com/programmierigel/pwmanager/enviornment"
	"github.com/programmierigel/pwmanager/storage/inmemory"
)

func main() {
	port, err := enviornment.Port(3000)
	if err != nil {
		panic(err)
	}

	path := enviornment.Path(".")
	password := enviornment.Password("123")

	store := inmemory.New(path, password)

	router := api.GetRouter(store)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
