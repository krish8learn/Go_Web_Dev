package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("or.gohtml"))
}

func main() {
	g1 := struct {
		Score1 int
		Score2 int
	}{
		4,
		5,
	}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatal(err)
	}
}
