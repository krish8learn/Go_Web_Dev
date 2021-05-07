package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0], " ", os.Args[1])

	str := fmt.Sprintf(`<!DOCTYPE html>
	<html lang = "en">
	<head>
	<meta charset = "UTF-8">
	<title>Hello world</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
`)

	//creating new file
	file1, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error while creating file")
	}

	io.Copy(file1, strings.NewReader(str))
}
