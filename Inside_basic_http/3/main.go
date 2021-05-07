package main

import (
	"html/template"
	"log"
	"net/http"
)

type hand int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("source.gohtml"))
}

func main() {
	var ha hand
	http.ListenAndServe(":8080", ha)
}

func (ha hand) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "source.gohtml", r.Form)

}
