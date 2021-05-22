package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/test1", test1)
	http.Handle("/test2", http.HandlerFunc(test2))
	http.ListenAndServe(":8080", nil)
}

func test1(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<h1>TABLE CREATED</h1><p>click below to go to WELCOME PAGE</p><a href="/welcome">go to welcome</a>`)
}

//<h2>CREATED TABLE</h2><br><p><strong>click the link below for WELCOME page</strong></p>
//	<h1><a href="/welcome">WELCOME</a></h1>
//<p>torres</p><h1><a href="/test2">test2</a></h1>
func test2(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `<h1>test2</h1>`)
}
