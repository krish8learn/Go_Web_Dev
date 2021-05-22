package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

var dsn string = "root:Krish@knight8@/Example"

var tpl *template.Template

type Person struct {
	UserName  string
	FirstName string
	LastName  string
	PassWord  string
}

type PersonUpdate struct {
	UserName string
	PassWord string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	DB, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalln("Error while connecting to the mysql DB:-->", err)
	}

	db = DB
	err = db.Ping()

	if err != nil {
		log.Fatalln("Error while checking status connection:-->", err)
	}

	fmt.Println("Database connection status all right")
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/createDB", createDB)
	http.HandleFunc("/insertDB", insertDB)
	http.HandleFunc("/updateDB", updateDB)
	http.HandleFunc("/selectDB", selectDB)
	http.HandleFunc("/deleteDB", deleteDB)
	http.HandleFunc("/selectAllDB", selectAllDB)
	http.ListenAndServe(":8080", nil)
}

func welcome(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "welcome.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func selectDB(w http.ResponseWriter, req *http.Request) {
	var getperson PersonUpdate
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")

		getperson, err := getUser(username, password, w)
		checkErr(err, w)

		str := fmt.Sprintf(`<h1>FirstName : %s, LastName: %s </h1> `, getperson.FirstName, getperson.LastName)

		fmt.Fprintln(w, str)
	}
	err := tpl.ExecuteTemplate(w, "select.gohtml", getperson)
	checkErr(err, w)
}

func insertDB(w http.ResponseWriter, req *http.Request) {
	var dataperson Person
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		dataperson = Person{username, firstname, lastname, password}

		stmt, err := db.Prepare(`INSERT INTO example (UserName, FirstName, LastName, PassWord) VALUES (?,?,?,?)`)
		checkErr(err, w)
		defer stmt.Close()
		_, err = stmt.Exec(dataperson.UserName, dataperson.FirstName, dataperson.LastName, dataperson.PassWord)
		checkErr(err, w)

		fmt.Fprintln(w, `<h1>Insertion Successful</h1>`)
	}

	err := tpl.ExecuteTemplate(w, "insert.gohtml", dataperson)
	checkErr(err, w)
}

func updateDB(w http.ResponseWriter, req *http.Request) {
	var update PersonUpdate
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")
		update = PersonUpdate{username, password}
		query := `UPDATE example SET PassWord = ? WHERE UserName = ?;`

		stmt, err := db.Prepare(query)
		checkErr(err, w)
		defer stmt.Close()

		_, err = stmt.Exec(username, password)
		checkErr(err, w)
		fmt.Fprintln(w, `<h1>Update successfull</h1>`)
	}
	err := tpl.ExecuteTemplate(w, "update.gohtml", update)
	checkErr(err, w)
}

func deleteDB(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		query := `DELETE FROM example WHERE UserName = ?;`

		stmt, err := db.Prepare(query)
		checkErr(err, w)
		_, err = stmt.Exec(username)
		checkErr(err, w)
		io.WriteString(w, `<h1>DELETION successful</h1>`)
	}
	err := tpl.ExecuteTemplate(w, "delete.gohtml", nil)
	checkErr(err, w)
}

func selectAllDB(w http.ResponseWriter, req *http.Request) {
	var listperson []Person

	query := "SELECT UserName, FirstName, LastName, PassWord FROM example;"

	rows, err := db.Query(query)
	defer rows.Close()
	checkErr(err, w)
	for rows.Next() {
		var dataperson Person
		err = rows.Scan(&dataperson.UserName, &dataperson.FirstName, &dataperson.LastName, &dataperson.PassWord)
		checkErr(err, w)
		listperson = append(listperson, dataperson)
	}

	err = tpl.ExecuteTemplate(w, "selectall.gohtml", listperson)
	checkErr(err, w)
}

func dropDB(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		tablename := req.FormValue("tablename")

		stmt, err := db.Prepare(`DROP TABLE ?;`)
		checkErr(err, w)
		defer stmt.Close()
		_, err = stmt.Exec(tablename)
		checkErr(err, w)

		io.WriteString(w, `<h1>Table Droped</h1><br><p><strong>click the link below for WELCOME page</strong></p>
		<h1><a href="/welcome">WELCOME</a></h1>`)
	}
}

func createDB(w http.ResponseWriter, req *http.Request) {
	createQuery := "CREATE TABLE example (example_id INT NOT NULL AUTO_INCREMENT,UserName VARCHAR(255) NOT NULL UNIQUE,FirstName VARCHAR(255), LastName VARCHAR(255) DEFAULT NULL, PassWord VARCHAR(255) DEFAULT NULL, Primary key(example_id));"
	//preparing the query
	stmt, err := db.Prepare(createQuery)
	checkErr(err, w)
	defer stmt.Close()
	//executing query
	_, err = stmt.Exec()
	checkErr(err, w)

	fmt.Fprintln(w, `<h1>TABLE CREATED</h1><h1><a href="/welcome">click this link to go to WELCOME PAGE</a></h1>`)

}

func checkErr(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func getUser(username, password string, w http.ResponseWriter) (*Person, error) {
	query := `SELECT FirstName, LastName FROM example WHERE UserName = ? AND PassWord = ?;`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var getperson Person
	row := stmt.QueryRow(username, password)
	err = row.Scan(&getperson.FirstName, &getperson.LastName)
	if err != nil {
		return nil, err
	}

	getperson.UserName = username
	getperson.PassWord = password

	return &getperson, nil
}

//query which need input --> Prepare() if that returns single row then ---> QueryRow()
//drop query will be Prepare() then Exec()
//query which returns multipe rows without any input Query() --- then Next() ---- then Scan()
