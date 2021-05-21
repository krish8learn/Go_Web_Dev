package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}
func main() {
	http.HandleFunc("/", foo)
	http.Handle("favicon.ico/", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)

	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	PrintGivenData := string(bs)

	err := tpl.ExecuteTemplate(w, "index.gohtml", PrintGivenData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
