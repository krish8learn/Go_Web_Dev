package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	fmt.Fprintln(w, "Got:-" +v)
}

/*here we pass data using url http://localhost:8080/?q=value*/

