package main

import (
	"log"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", firstpage)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/registrationS", registrationS)
	http.HandleFunc("/data", data)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

type Person struct {
	UserName  string
	Password  []byte
	FirstName string
	LastName  string
}

var DBperson = make(map[string]Person)

func firstpage(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "firstpage.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("Error while executing firstpage template: ", err)
	}
}

func signup(w http.ResponseWriter, req *http.Request) {

	var person Person
	if req.Method == http.MethodPost {
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		username := req.FormValue("username")
		password := req.FormValue("password")

		bspassword, err := bcrypt.GenerateFromPassword([]byte(password), 1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln("Error while encrypting password: ", err)
			return
		}

		person = Person{username, bspassword, firstname, lastname}

		DBperson[username] = person

		cookie := &http.Cookie{
			Name:  "Registration_Session_Cookie",
			Value: username,
		}

		http.SetCookie(w, cookie)

		http.Redirect(w, req, "/registrationS", http.StatusSeeOther)
	}

	err := tpl.ExecuteTemplate(w, "signup.gohtml", person)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("Error while executing signup template: ", err)
	}
}

func registrationS(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("Registration_Session_Cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Fatalln("Error while receving registration cookie: ", err)
	}

	RegiPerson := DBperson[cookie.Value]

	err = tpl.ExecuteTemplate(w, "registrationS.gohtml", RegiPerson)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("Error while executing Registration template:", err)
	}
}

func login(w http.ResponseWriter, req *http.Request) {

}

func data(w http.ResponseWriter, req *http.Request) {

}
