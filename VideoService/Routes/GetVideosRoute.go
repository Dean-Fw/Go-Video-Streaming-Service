package routes

import (
	"net/http"
	"videoservice/Interfaces"
)

type GetVideosRoute struct {
	Handler interfaces.IHandler
}

func (gv GetVideosRoute) Add(router *http.ServeMux) {
	VideosHandler := gv.Handler

	router.HandleFunc("GET /videos/{name}", VideosHandler.Handle)
}
