package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/verify", verify)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func welcome(w http.ResponseWriter, req *http.Request) {
	//checking cookie present or not
	cookie, err := req.Cookie("examplecookie")
	//creating the cookie if not present
	if err != nil {
		cookie = &http.Cookie{
			Name:  "examplecookie",
			Value: "",
		}
	}

	//getting the email from client and putting it inside cookie.value
	if req.Method == http.MethodPost {
		email := req.FormValue("email")
		cookie.Value = email + `|` + getcode(email)
	}

	//setting the cookie
	http.SetCookie(w, cookie)

	//page output html
	io.WriteString(w, `<!DOCTYPE html>
	<html>
	  <body>
	    <form method="POST">
	      <input type="email" name="email">
	      <input type="submit">
	    </form>
	    <a href="/verify">Validate This `+cookie.Value+`</a>
	  </body>
	</html>`)

}

func verify(w http.ResponseWriter, req *http.Request) {
	//chekcing the cookie
	cookie, err := req.Cookie("examplecookie")
	if err != nil {
		http.Redirect(w, req, "/welcome", http.StatusSeeOther)
		return
	}

	if cookie.Value == "" {
		http.Redirect(w, req, "/welcome", http.StatusSeeOther)
		return
	}

	xs := strings.Split(cookie.Value, "|")

	email := xs[0]
	codereceived := xs[1]
	checkcode := getcode(email)

	if codereceived != checkcode {
		fmt.Println("HAMC code didnt match")
		fmt.Println(codereceived)
		fmt.Println(checkcode)
		http.Redirect(w, req, "/welcome", http.StatusSeeOther)
		return
	}

	io.WriteString(w, `<!DOCTYPE html>
	<html>
	  <body>
	  	<h1>`+codereceived+` - RECEIVED </h1>
	  	<h1>`+checkcode+` - RECALCULATED </h1>
	  </body>
	</html>`)

}

func getcode(data string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func checkErr(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
