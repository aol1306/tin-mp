package main

import (
	"html/template"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)
	} else {
		// TODO: verify login credentials
		http.Redirect(w, r, "landing", http.StatusSeeOther)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))
	tmpl.Execute(w, nil)
}

func viewcards(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/view-cards.html"))
	tmpl.Execute(w, nil)
}

// TODO: add auth

func landing(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/landing.html"))
	tmpl.Execute(w, nil)
}

func edit(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/edit.html"))
	tmpl.Execute(w, nil)
}

func details(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/details.html"))
	tmpl.Execute(w, nil)
}

func add(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/add.html"))
	tmpl.Execute(w, nil)
}

func main() {
	log.Println("Start")

	http.HandleFunc("/", logging(index))
	http.HandleFunc("/login", logging(login))
	http.HandleFunc("/register", logging(register))
	http.HandleFunc("/viewcards", logging(viewcards))
	http.HandleFunc("/landing", logging(landing))
	http.HandleFunc("/edit", logging(edit))
	http.HandleFunc("/details", logging(details))
	http.HandleFunc("/add", logging(add))

	// static files
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
