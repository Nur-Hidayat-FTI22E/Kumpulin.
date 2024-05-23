package main

import (
	"koding/config"
	"koding/controllers"
	"net/http"
)

func main() {
	config.Connection()

	http.HandleFunc("/", controllers.CreateTugas)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":3000", nil)
}
