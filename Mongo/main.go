package main

import (
	"net/http"

	"github.com/krish8learn/Go_Web_Dev/Mongo/controller"
)

func main() {
	http.HandleFunc("/welcome", controller.Welcome)
	http.HandleFunc("/create", controller.CreateOne)
	http.HandleFunc("/readall", controller.ReadAll)
	http.HandleFunc("/readone", controller.ReadOne)
	http.HandleFunc("/update", controller.Update)
	http.HandleFunc("/delete", controller.Delete)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
