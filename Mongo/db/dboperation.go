package db

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/krish8learn/Go_Web_Dev/Mongo/model"
	"gopkg.in/mgo.v2/bson"
)

func CreateDB(r *http.Request) error {

	var RecvBook model.Book
	RecvBook.Isbn = r.FormValue("isbn")
	RecvBook.Title = r.FormValue("title")
	RecvBook.Author = r.FormValue("author")
	p, err := strconv.ParseFloat(r.FormValue("price"), 32)
	if err != nil {
		return errors.New("Must be number")
	}
	RecvBook.Price = float32(p)

	if RecvBook.Isbn == "" || RecvBook.Title == "" || RecvBook.Author == "" {
		return errors.New("Fill all fields completely")
	}

	err = Clsn.Insert(RecvBook)
	if err != nil {
		return errors.New("500 internal server error")
	}

	return nil
}

func ReadAllDB() ([]model.Book, error) {
	var allbook []model.Book
	err := Clsn.Find(bson.M{}).All(&allbook)
	if err != nil {
		return nil,err
	}
	return allbook, nil
}

func ReadOne(author string) (*model.Book, error) {
	var OneBook model.Book
	err := Clsn.Find(bson.M{"author":author}).One(&OneBook)

	if err != nil {
		return nil,errors.New("Fille with Correct information")
	}
	return &OneBook, nil
}

func UpdateDB(req *http.Request) ( error) {
	title := req.FormValue("title")
	author := req.FormValue("author")
	if title == ""||author ==""{
		return errors.New("Do not enter blank fileds")
	}

	err := Clsn.Update(bson.M{"title":title},bson.M{"author":author})

	if err != nil {
		return errors.New("Give right title information present into DB")
	}

	return nil
}

func DeleteDB(title string) error {
	err := Clsn.Remove(bson.M{"title":title})

	if err != nil {
		return errors.New("Enter proper data present in DB")
	}
	return nil
}
