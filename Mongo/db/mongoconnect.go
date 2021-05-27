package db

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
)

var DB *mgo.Database

var Clsn *mgo.Collection

var mongologin = "mongodb://bond:moneypenny007@localhost/bookstore"

func init() {
	session, err := mgo.Dial(mongologin)

	if err != nil {
		log.Fatalln(err)
	}

	if session.Ping(); err != nil {
		log.Fatalln(err)
	}

	DB = session.DB("bookstore")

	Clsn = DB.C("books")

	fmt.Println("Mongo DB connection Normal")
}
