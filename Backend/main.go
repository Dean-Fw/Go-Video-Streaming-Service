package main

import (
	"html/template"
	"log"
	"net/http"
	"videoservice/Controllers"
	"videoservice/Services"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("../Templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(w, nil)
}

func main() {

	videoService := controllers.VideoController{
		FileSystemService: services.FileSystemService{},
	}

	http.HandleFunc("GET /video/{name}", videoService.GetVideos)

	http.HandleFunc("GET /{$}", homePageHandler)

	log.Print("[SERVER] : Running on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
