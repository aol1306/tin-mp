package main

import (
	"html/template"
	"log"
	"net/http"
)

/*func logging(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.URL.Path)
        f(w, r)
    }
}*/

func main() {
    log.Println("Start")

	// index page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	// login page
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl := template.Must(template.ParseFiles("templates/login.html"))
			tmpl.Execute(w, nil)
		} else {
			// TODO: verify login credentials
			http.Redirect(w, r, "landing", http.StatusSeeOther)
		}
	})

	// register page
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/register.html"))
		tmpl.Execute(w, nil)
	})

    // ***
    // ONLY FOR AUTHORIZED USERS
    // ***

    // landing page
    http.HandleFunc("/landing", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/landing.html"))
		tmpl.Execute(w, nil)
	})

    // edit page
    http.HandleFunc("/edit", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/edit.html"))
		tmpl.Execute(w, nil)
	})

    // details page
    http.HandleFunc("/details", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/details.html"))
		tmpl.Execute(w, nil)
	})

	// static files
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
