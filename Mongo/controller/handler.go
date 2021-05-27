package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/krish8learn/Go_Web_Dev/Mongo/db"
	"github.com/krish8learn/Go_Web_Dev/Mongo/tpl"
)

func Welcome(w http.ResponseWriter, req *http.Request) {
	err := tpl.TPL.ExecuteTemplate(w, "welcome.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("Error in WELCOME page :-->", err)
	}
}

func CreateOne(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := db.CreateDB(r)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			log.Println("Error CreateOne handler", err)
		}
		fmt.Fprintln(w, `<h1>DATA CREATED</h1><h1><a href="/welcome">click this link to go to WELCOME PAGE</a></h1>`)
	}

	err := tpl.TPL.ExecuteTemplate(w, "create.gohtml", nil)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		log.Println("Error in ReadAll handlers, template execution : ", err)
		return
	}

}

func ReadAll(w http.ResponseWriter, req *http.Request) {
	alldata, err := db.ReadAllDB()

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		log.Println("Error in ReadAll handler: ", err)
		return
	}

	err = tpl.TPL.ExecuteTemplate(w, "readall.gohtml", alldata)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		log.Println("Error in ReadAll handlers, template execution : ", err)
		return
	}
}

/*func ReadOne1(w http.ResponseWriter, req *http.Request) {
	err := tpl.TPL.ExecuteTemplate(w, "readone1.gohtml", nil)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		log.Println("Error in ReadOne1 handler while template parsing")
	}
}*/
func ReadOne(w http.ResponseWriter, req *http.Request) {
	//var b model.Book
	if req.Method == http.MethodGet {

		err := tpl.TPL.ExecuteTemplate(w, "readone1.gohtml", nil)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			log.Println("Error in ReadOne2 Handler, while parsing executing template", err)
			return
		}

		author := req.FormValue("author")
		if author == "" {
			http.Error(w, "Fill the required field", http.StatusBadRequest)
			return
		}
		//var err error
		book, err := db.ReadOne(author)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			log.Println("Error in ReadOne Handler, while receiving book data", err)
			return
		}

		err = tpl.TPL.ExecuteTemplate(w, "readone.gohtml", *book)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			log.Println("Error in ReadOne Handler data process, while parsing executing template", err)
			return
		}

	}

	//fmt.Println(w, "Read successful", *book)

}

func Update(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		err := db.UpdateDB(req)
		if err != nil {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			log.Println("Error in Update Handler, bad request", err)
			return
		}
		fmt.Fprintln(w, `<p>UPDATED</p><br><p><a href="/welcome">GO TO WELCOME</a></p>`)
	}

	err := tpl.TPL.ExecuteTemplate(w, "update.gohtml", nil)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		log.Println("Error in Update Handler, while parsing executing template", err)
		return
	}
}

func Delete(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		err := tpl.TPL.ExecuteTemplate(w, "delete.gohtml", nil)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			log.Println("Error in Delete Handler, while parsing executing template", err)
			return
		}

		book := req.FormValue("title")
		if book == "" {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
		}
		if err := db.DeleteDB(book); err != nil {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			log.Println("Error in DELETE Handler, bad request", err)
			return
		}
		io.WriteString(w, `<p>DELETED</p><br><p><a href="/welcome">GO TO WELCOME</a></p>`)
	}

}
