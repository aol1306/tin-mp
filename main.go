package main

import (
	"html/template"
	"log"
	"net/http"
	"tinmp/model"

	"github.com/gobuffalo/packr/v2"
)

var tmplBox = packr.New("tmpl", "./templates")
var staticBox = packr.New("static", "./static")

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	s, err := tmplBox.FindString("index.html")
	if err != nil {
		log.Println(err)
	}
	tmpl, err := template.New("index").Parse(s)
	if err != nil {
		log.Println(err)
	}
	tmpl.Execute(w, nil)
}

type LoginTemplateData struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s, _ := tmplBox.FindString("login.html")
		tmpl, _ := template.New("login").Parse(s)
		tmpl.Execute(w, LoginTemplateData{Message: ""})
	} else {
		if model.VerifyUser(r.FormValue("username"), r.FormValue("password")) {
			// TODO: authorize
			http.Redirect(w, r, "landing", http.StatusSeeOther)
		} else {
			s, _ := tmplBox.FindString("login.html")
			tmpl, _ := template.New("login").Parse(s)
			tmpl.Execute(w, LoginTemplateData{Message: "Nieprawidłowy login lub hasło!"})
		}
	}
}

type RegisterTemplateData struct {
	Message string
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s, _ := tmplBox.FindString("register.html")
		tmpl, _ := template.New("register").Parse(s)
		tmpl.Execute(w, RegisterTemplateData{Message: ""})
	} else {
		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")
		passwordConfirm := r.FormValue("confirm-password")
		log.Println(password)
		log.Println(passwordConfirm)
		// validate
		errMsg := ""
		if email == "" || username == "" || password == "" || passwordConfirm == "" {
			errMsg = "Nie wszystkie pola zostały wypełnione"
		}
		if password != passwordConfirm {
			errMsg = "Hasła nie są identyczne"
		}
		if len(password) < 8 {
			errMsg = "Hasło jest za krótkie"
		}
		if errMsg != "" {
			s, _ := tmplBox.FindString("register.html")
			tmpl, _ := template.New("register").Parse(s)
			tmpl.Execute(w, RegisterTemplateData{Message: errMsg})
		} else {
			// perform registration
			model.RegisterUser(username, email, password)
			http.Redirect(w, r, "login", http.StatusSeeOther)
		}
	}
}

func viewcards(w http.ResponseWriter, r *http.Request) {
	s, _ := tmplBox.FindString("view-cards.html")
	tmpl, _ := template.New("view-cards").Parse(s)
	tmpl.Execute(w, nil)
}

// TODO: add auth

func landing(w http.ResponseWriter, r *http.Request) {
	s, _ := tmplBox.FindString("landing.html")
	tmpl, _ := template.New("landing").Parse(s)
	tmpl.Execute(w, nil)
}

func edit(w http.ResponseWriter, r *http.Request) {
	s, _ := tmplBox.FindString("edit.html")
	tmpl, _ := template.New("edit").Parse(s)
	tmpl.Execute(w, nil)
}

func details(w http.ResponseWriter, r *http.Request) {
	s, _ := tmplBox.FindString("details.html")
	tmpl, _ := template.New("details").Parse(s)
	tmpl.Execute(w, nil)
}

func add(w http.ResponseWriter, r *http.Request) {
	s, _ := tmplBox.FindString("add.html")
	tmpl, _ := template.New("add").Parse(s)
	tmpl.Execute(w, nil)
}

func main() {
	log.Println("Start")

	model.Init()

	http.HandleFunc("/", logging(index))
	http.HandleFunc("/login", logging(login))
	http.HandleFunc("/register", logging(register))
	http.HandleFunc("/viewcards", logging(viewcards))
	http.HandleFunc("/landing", logging(landing))
	http.HandleFunc("/edit", logging(edit))
	http.HandleFunc("/details", logging(details))
	http.HandleFunc("/add", logging(add))

	// static files
	fs := http.FileServer(staticBox)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
