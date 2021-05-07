package main

import "fmt"

func main(){
	name := "Krish"

	template := `<!DOCTYPE html>
				<html lang = "en">
				<head>
				<meta charset = "UTF-8">
				<title>Hello world</title>
				</head>
				<body>
				<h1>` + name + `</h1>
				</body>
				</html>`

	fmt.Println(template)			
}