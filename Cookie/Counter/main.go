package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/Counter", CounterCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func CounterCookie(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("Counter_Cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "Counter_Cookie",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln("Atoi", err)
	}
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)
	//fmt.Fprintln(w, cookie.Value)
}
