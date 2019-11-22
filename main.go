package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
    // index page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
