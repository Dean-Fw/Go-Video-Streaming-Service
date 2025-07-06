package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type GetHomePageHandler struct{}

func (GetHomePageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("../Templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(w, nil)
}
