package main

import (
	"log"
	"os"
	"text/template"
)

func double(z int) int {
	return 2 * z
}

func square(z int) int {
	return z * z
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

var fm = template.FuncMap{
	"double": double,
	"square": square,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 9)
	if err != nil {
		log.Fatalln(err)
	}

	file3, err3 := os.Create("index3.html")
	if err3 != nil {
		log.Fatal(err3)
	}

	exeErr := tpl.ExecuteTemplate(file3, "tpl.gohtml", 9)
	if exeErr != nil {
		log.Fatalln(err)
	}
}
