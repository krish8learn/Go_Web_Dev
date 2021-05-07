package main

import (
	"log"
	"os"
	"text/template"
)

type player struct {
	Name   string
	Number int
}

type club struct {
	Cname    string
	Location string
}

type items struct {
	Person []player
	Object []club
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("Untitled-1"))
}
func main() {
	a := player{
		Name:   "Mbappe",
		Number: 9,
	}
	b := player{
		Name:   "Neymar",
		Number: 11,
	}

	c := club{
		Cname:    "PSG",
		Location: "PARIS",
	}

	d := club{
		Cname:    "PSG",
		Location: "PARIS",
	}

	playerlist := []player{a, b}
	clublist := []club{c, d}

	item := items{
		Person: playerlist,
		Object: clublist,
	}

	file, flieErr := os.Create("New.html")
	if flieErr != nil {
		log.Fatal(flieErr)
	}

	FileExeErr := tpl.Execute(file, item)
	if FileExeErr != nil {
		log.Fatal(FileExeErr)
	}

	outputErr := tpl.Execute(os.Stdout, item)
	if outputErr != nil {
		log.Fatal(outputErr)
	}
}
