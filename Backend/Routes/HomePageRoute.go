package routes

import (
	"net/http"
	"videoservice/Interfaces"
)

type GetHomePageRoute struct {
	Handler interfaces.IHandler
}

func (gh GetHomePageRoute) Add(router *http.ServeMux) {
	router.HandleFunc("GET /{$}", gh.Handler.Handle)
}
