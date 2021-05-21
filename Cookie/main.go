package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/set", SettingCookie)
	http.HandleFunc("/read", readingCookie)
	http.Handle("favicon.ico/", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func SettingCookie(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "First_Cookie",
		Value: "Nothing",
	})

	fmt.Fprintln(w, "Cookie Set")
	fmt.Fprintf(w, "chrome-->dev tools--->application--->cookies")

}

func readingCookie(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("First_Cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	fmt.Fprintln(w, "COOKIE----", c)
	fmt.Println(c)
}
