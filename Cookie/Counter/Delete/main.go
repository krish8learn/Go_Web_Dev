package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<h1><a href="/set">set cookie</h1>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "Biscuits",
		Value: "Oreo",
	})
	io.WriteString(w, `<h1>Set page</h1>`)
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Biscuits")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintln(w, "COOKIE NAME AND VALUE", c.Name, c.Value)

	fmt.Fprintln(w, `<h1><a href="/expire">expire cookie</h1>`)
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Biscuits")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	c.MaxAge = -1 //-1 ---> delete cookie

	http.SetCookie(w, c)
	http.Redirect(w, req, "/set", http.StatusSeeOther)
}
