package main

import (
	"net/http"
	"videoservice/Handlers"
	"videoservice/Routes"
	"videoservice/Services"
)

func AddRoutes() *http.ServeMux {
	router := http.NewServeMux()

	AddGetHomePage(router)
	AddGetVideos(router)

	return router
}

func AddGetHomePage(router *http.ServeMux) {
	GetHomePageRoute := routes.GetHomePageRoute{
		Handler: handlers.GetHomePageHandler{},
	}

	GetHomePageRoute.Add(router)
}

func AddGetVideos(router *http.ServeMux) {
	GetVideosHandler := handlers.GetVideosHandler{
		FileSystemService: services.FileSystemService{},
	}

	GetVideosRoute := routes.GetVideosRoute{
		Handler: GetVideosHandler,
	}

	GetVideosRoute.Add(router)
}
