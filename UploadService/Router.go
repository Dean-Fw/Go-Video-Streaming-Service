package main

import (
	"fmt"
	"net/http"
	"uploadservice/Handlers"
	"uploadservice/Services"
)

func AddRoutes() *http.ServeMux {
	router := http.NewServeMux()

	postStartUploadHandler := handlers.PostStartUploadHandler{
		FileSystemService: services.FileSystemService{},
		HashingService:    services.HashingService{},
	}

	corsHandler := handlers.CorsHandler{}

	router.HandleFunc("POST /videos", postStartUploadHandler.Handle)
	router.HandleFunc("OPTIONS /videos", corsHandler.Handle)
	router.HandleFunc("GET /videos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})

	return router
}
