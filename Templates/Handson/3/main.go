package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

type Region struct {
	Hotels []Hotel
	Area   string
}

func main() {
	input := []Region{
		{
			Hotels: []Hotel{
				{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
				},
				{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
				},
			},
			Area: "Southern",
		},
		{
			Hotels: []Hotel{
				{
					Name:    "Hotel Veas",
					Address: "42 Sunset Angelas",
					City:    "Los Vegas",
					Zip:     "956134",
				},
				{
					Name:    "k",
					Address: "3",
					City:    "J",
					Zip:     "94612",
				},
			},
			Area: "Northern",
		},
	}

	err := tpl.Execute(os.Stdout, input)
	if err != nil {
		log.Fatal(err)
	}
}
