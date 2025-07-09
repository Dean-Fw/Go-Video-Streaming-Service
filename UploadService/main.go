package main

import (
	"log"
	"net/http"
)

func main() {
	router := AddRoutes()

	server := http.Server{
		Addr:    ":5050",
		Handler: router,
	}

	log.Print("[SERVER] : Running on Port 5050")
	log.Fatal(server.ListenAndServe())
}
