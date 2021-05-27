package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/krish8learn/Go_Web_Dev/MVC/controller"
	"github.com/krish8learn/Go_Web_Dev/MVC/model"
)

func main() {
	router := httprouter.New()
	cc := controller.NewUserController(getSession())
	router.GET("/", cc.Getuser)
	router.POST("/post", cc.CreateUser)
	router.DELETE("/delete", cc.Deleteuser)
	http.ListenAndServe(":8080", router)
}

func getSession() map[string]model.User {
	return make(map[string]model.User)
}
