package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/marsh", marsh)
	http.HandleFunc("/enc", enc)
	http.ListenAndServe(":8080", nil)
}

type Person struct {
	Fname string
	Lname string
	Item  []string
}

func foo(w http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
			<html lang = "en">
			<head>
			<meta charset="UTF-8">
			<title>FOO</title>
			</head>
			<body>foo</body>
			</html>`
	w.Write([]byte(s))
}

func marsh(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := Person{
		Fname: "krish",
		Lname: "knight",
		Item:  []string{"food1", "food2"},
	}

	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func enc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := Person{
		Fname: "krish",
		Lname: "knight",
		Item:  []string{"food1", "food2"},
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Fatalln(err)
	}
}
