package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl, tpl2, tpl3, tpl4 *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tpl2 = template.Must(template.ParseFiles("tpl2.gohtml"))
	tpl3 = template.Must(template.ParseFiles("tpl3.gohtml"))
	tpl4 = template.Must(template.ParseFiles("tpl4.gohtml"))

}

type player struct {
	Name   string
	Number int
}

func main() {
	lists := []string{"Gandhi", "kane", "torres", "liverpool"}
	maps := map[string]string{"liverpool": "torres", "chelsea": "drogba", "Man utd": "ronaldo"}
	torres := player{
		Name:   "Fernando Torres",
		Number: 9,
	}
	structlist := []player{
		player{
			Name : "Drogba",
			Number: 11,
		},
		player{
			Name: "Lampard",
			Number: 8,
		},
	}

	file1, err1 := os.Create("index.html")
	if err1 != nil {
		log.Fatal(err1)
	}

	file2, err2 := os.Create("index2.html")
	if err2 != nil {
		log.Fatal(err2)
	}

	file3 , err3 := os.Create("index3.html")
	if err3 != nil{
		log.Fatal(err3)
	}

	file4, err4 := os.Create("index4.html")
	if err4 != nil{
		log.Fatal(err4)
	}

	errExe1 := tpl.Execute(file1, lists)
	if errExe1 != nil {
		log.Fatal(errExe1)
	}

	errExe2 := tpl2.Execute(file2, maps)
	if errExe2 != nil {
		log.Fatal(errExe2)
	}

	errExe3 := tpl3.Execute(file3, torres)
	if errExe3 != nil {
		log.Fatal(errExe3)
	}

	errEx4 := tpl4.Execute(file4,structlist)
	if errEx4 != nil{
		log.Fatal(errEx4)
	}

	fmt.Println("List")
	errOutput1 := tpl.Execute(os.Stdout, lists)
	if errOutput1 != nil {
		log.Fatal(errOutput1)
	}

	fmt.Println("maps")
	errOutput2 := tpl2.Execute(os.Stdout, maps)
	if errOutput2 != nil {
		log.Fatal(errOutput2)
	}

	fmt.Println("struct")
	errOutput3 := tpl3.Execute(os.Stdout, torres)
	if errOutput3 != nil {
		log.Fatal(errOutput3)
	}

	fmt.Println("StructList")
	errOutput4:= tpl4.Execute(os.Stdout, structlist)
	if errOutput4 != nil{
		log.Fatal(errOutput4)
	}

}
