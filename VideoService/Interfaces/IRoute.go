package interfaces

import "net/http"

type IRoute interface {
	Add(router *http.ServeMux)
}
