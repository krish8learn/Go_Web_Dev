package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	//fs := http.FileServer(http.Dir("DogPic"))
	http.Handle("/resources/",http.StripPrefix("/resources",http.FileServer(http.Dir("DogPic"))) )
	http.Handle("/", http.HandlerFunc(dog))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("Templates/Index.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
