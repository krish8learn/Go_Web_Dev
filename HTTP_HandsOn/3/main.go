package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/bar/", http.HandlerFunc(bar))
	http.Handle("/me/", http.HandlerFunc(mcleod))

	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func bar(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bar ran")
}

func mcleod(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")

	if err != nil {
		log.Fatalln("Error while parsing template")
	}

	err = tpl.ExecuteTemplate(w, "something.gohtml", "krish")
	if err != nil {
		log.Fatalln("Error while executing file")
	}
}
