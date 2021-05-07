package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hand string

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("source.gohtml"))
}

func (ha hand) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method     string
		Submission url.Values
	}{
		r.Method,
		r.Form,
	}

	tpl.ExecuteTemplate(w, "source.gohtml", data)
}

func main() {
	var ha hand
	http.ListenAndServe(":8082", ha)
}
