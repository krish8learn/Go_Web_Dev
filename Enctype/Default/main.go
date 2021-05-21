package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template/index.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("favicon.ico/", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

type Person struct {
	FirstName string
	LastName  string
	Subscribe bool
}

func foo(w http.ResponseWriter, req *http.Request) {
	//reading the data given by client
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	PrintGivenData := string(bs)

	err := tpl.ExecuteTemplate(w, "index.gohtml", PrintGivenData)

	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
