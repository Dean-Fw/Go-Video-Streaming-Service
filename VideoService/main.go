package main

import (
	"log"
	"net/http"
)

func main() {
	router := AddRoutes()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Print("[SERVER] : Running on Port 8080")
	log.Fatal(server.ListenAndServe())
}
