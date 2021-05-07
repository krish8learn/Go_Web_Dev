package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main(){
	lists := []string{"krish", "kane", "kevin"}

	err := tpl.Execute(os.Stdout, lists)
	if err!= nil{
		log.Fatal(err)
	}

	P := struct{
		Words []string
		lname  string
	}{
		lists,
		"krish",
	}

	err = tpl.Execute(os.Stdout, P)
	if err != nil{
		log.Fatal(err)
	}
}