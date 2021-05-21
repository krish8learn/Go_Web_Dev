package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

//here we get value using POST in the body/payload
/*func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<form method = "POST">
	<input type="text" name="q">
	<input type = "submit">
	</form>
	<br>`+v)
}*/
// such as http://localhost:8080

//here we get value using GET in the url
func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<form method = "GET">
	<input type="text" name="q">
	<input type = "submit">
	</form>
	<br>`+v)
}

// such as http://localhost:8080/?q=friends+season+4
