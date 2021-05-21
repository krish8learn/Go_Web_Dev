package main

import (
	"log"
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

type Person struct {
	UserName  string
	FirstName string
	LastName  string
	Subscribe bool
}

var DBsession = make(map[string]string) //sessionid , username
var DBperson = map[string]Person{}      //username, person

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {

	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicone.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	//search for cookie, if not found then generate cookie with id
	cookie, err := req.Cookie("Session")
	if err != nil {
		sessionID, err := uuid.NewV4()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln("Error in generating sessionID", err)
		}
		cookie = &http.Cookie{
			Name:  "Session",
			Value: sessionID.String(),
		}
	}
	http.SetCookie(w, cookie)

	//for existing
	var user Person
	username, cond := DBsession[cookie.Value]
	if cond == true {
		user = DBperson[username]
	}

	// new username
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		subscribe := req.FormValue("subscribe") == "on"
		user = Person{username, firstname, lastname, subscribe}
		DBsession[cookie.Value] = username
		DBperson[username] = user
	}

	tplerr := tpl.ExecuteTemplate(w, "set.gohtml", user)
	if tplerr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func read(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("Session")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	username, cond := DBsession[cookie.Value]
	if cond == false {
		http.Error(w, err.Error(), http.StatusNotFound)
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}
	user := DBperson[username]

	tplErr := tpl.ExecuteTemplate(w, "read.gohtml", user)
	if tplErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
