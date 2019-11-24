package controller

import (
	"html/template"
	"log"
	"net/http"
	"tinmp/model"

	"github.com/gobuffalo/packr/v2"
)

var tmplBox = packr.New("tmpl", "./../templates")

// Index /
func Index(w http.ResponseWriter, r *http.Request) {
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

type loginTemplateData struct {
	Message string
}

// Login /login
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s, _ := tmplBox.FindString("login.html")
		tmpl, _ := template.New("login").Parse(s)
		tmpl.Execute(w, loginTemplateData{Message: ""})
	} else {
		if model.VerifyUser(r.FormValue("username"), r.FormValue("password")) {
			// TODO: authorize
			http.Redirect(w, r, "landing", http.StatusSeeOther)
		} else {
			s, _ := tmplBox.FindString("login.html")
			tmpl, _ := template.New("login").Parse(s)
			tmpl.Execute(w, loginTemplateData{Message: "Nieprawidłowy login lub hasło!"})
		}
	}
}

type registerTemplateData struct {
	Message string
}

// Register /register
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s, _ := tmplBox.FindString("register.html")
		tmpl, _ := template.New("register").Parse(s)
		tmpl.Execute(w, registerTemplateData{Message: ""})
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
			tmpl.Execute(w, registerTemplateData{Message: errMsg})
		} else {
			// perform registration
			model.RegisterUser(username, email, password)
			http.Redirect(w, r, "login", http.StatusSeeOther)
		}
	}
}

// Viewcards /viewcards
func Viewcards(w http.ResponseWriter, r *http.Request) {
	s, _ := tmplBox.FindString("view-cards.html")
	tmpl, _ := template.New("view-cards").Parse(s)
	tmpl.Execute(w, nil)
}

// TODO: add auth

// Landing /landing
func Landing(w http.ResponseWriter, r *http.Request) {
	s, _ := tmplBox.FindString("landing.html")
	tmpl, _ := template.New("landing").Parse(s)
	tmpl.Execute(w, nil)
}

// Edit /edit
func Edit(w http.ResponseWriter, r *http.Request) {
	s, _ := tmplBox.FindString("edit.html")
	tmpl, _ := template.New("edit").Parse(s)
	tmpl.Execute(w, nil)
}

// Details /details
func Details(w http.ResponseWriter, r *http.Request) {
	s, _ := tmplBox.FindString("details.html")
	tmpl, _ := template.New("details").Parse(s)
	tmpl.Execute(w, nil)
}

// Add /add
func Add(w http.ResponseWriter, r *http.Request) {
	s, _ := tmplBox.FindString("add.html")
	tmpl, _ := template.New("add").Parse(s)
	tmpl.Execute(w, nil)
}
