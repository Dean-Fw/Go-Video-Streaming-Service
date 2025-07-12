package handlers

import (
	"fmt"
	"net/http"
)

type PatchVideosHandler struct{}

func (handler PatchVideosHandler) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")

}
