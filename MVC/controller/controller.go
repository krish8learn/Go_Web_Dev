package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/krish8learn/Go_Web_Dev/MVC/model"
)

type UserController struct {
	session map[string]model.User
}

func NewUserController(m map[string]model.User) *UserController {
	return &UserController{
		session: m,
	}
}

func (us UserController) Getuser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	input := p.ByName("id")

	ruser := us.session[input]

	bs, err := json.Marshal(ruser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	fmt.Fprintf(w, "%s\n", bs)
}

func (us UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	inputuser := model.User{}

	err := json.NewDecoder(r.Body).Decode(&inputuser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.NewEncoder(w).Encode(&inputuser)

	us.session[inputuser.ID] = inputuser

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	//fmt.Fprintf(w, "Created")
}

func (us UserController) Deleteuser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	input := p.ByName("id")
	delete(us.session, input)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "deleted")
}
