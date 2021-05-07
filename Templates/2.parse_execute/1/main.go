package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//taking data from the file
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	//creating an html file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}

	//executing wrting data of tpl into nf
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatal(err)
	}

	//executing the file command line output
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
}
