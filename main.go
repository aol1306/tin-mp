package main

import (
	"log"
	"net/http"
	"tinmp/controller"
	"tinmp/model"

	"github.com/gobuffalo/packr/v2"
)

var staticBox = packr.New("static", "./static")

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func main() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Start")

	model.Init()

	http.HandleFunc("/", logging(controller.Index))
	http.HandleFunc("/login", logging(controller.Login))
	http.HandleFunc("/register", logging(controller.Register))
	http.HandleFunc("/viewcards", logging(controller.Viewcards))
	http.HandleFunc("/landing", logging(controller.Landing))
	http.HandleFunc("/edit", logging(controller.Edit))
	http.HandleFunc("/details", logging(controller.Details))
	http.HandleFunc("/add", logging(controller.Add))

	// static files
	fs := http.FileServer(staticBox)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
