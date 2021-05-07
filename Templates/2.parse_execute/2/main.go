package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("temp/*"))
}

func main() {
	/*tpl, err := template.ParseGlob("temp/*")
	if err != nil {
		log.Fatal(err)
	}*/

	fmt.Println("Execution")
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Execution_template")
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
