package main

import (
	"net/http"
	"uploadservice/Handlers"
	"uploadservice/Services"
	transport "uploadservice/Transport"
)

func AddRoutes() *http.ServeMux {
	router := http.NewServeMux()

	addPostVideos(router)
	addPatchVideos(router)

	corsHandler := handlers.CorsHandler{}

	router.HandleFunc("OPTIONS /videos", corsHandler.Handle)

	return router
}

func addPostVideos(router *http.ServeMux) {
	postStartUploadHandler := handlers.PostStartUploadHandler{
		FileSystemService: services.FileSystemService{},
		HashingService:    services.HashingService{},
	}

	router.HandleFunc("POST /videos", postStartUploadHandler.Handle)
}

func addPatchVideos(router *http.ServeMux) {
	patchVideosHandler := handlers.PatchVideosHandler{
		FileSystemService: services.FileSystemService{},
	}

	patchVideosTranspot := transport.PatchVideosTransport{
		HeaderValidatorService: services.HeaderValidatorService{},
		PatchVideosHandler:     patchVideosHandler,
	}

	router.HandleFunc("PATCH /videos/{id}", patchVideosTranspot.Receive)
}
