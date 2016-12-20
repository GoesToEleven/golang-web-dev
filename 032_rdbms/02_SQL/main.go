package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "awsuser:mypassword@tcp(mydbinstance.cakwl95bxza0.us-west-1.rds.amazonaws.com:3306)/test02?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/retrieve", retrieve)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "at index")
	check(err)
}


func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT aName FROM amigos`)
	check(err)

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	_, err = fmt.Fprintln(w, "CREATED TABLE CUSTOMER", n)
	check(err)
}

func insert(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("James");`)
	check(err)

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	_, err = fmt.Fprintln(w, "INSERTED RECORD", n)
	check(err)
}

func retrieve(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer`)
	check(err)

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Println(name)

		_, err = fmt.Fprintln(w, "RETREIVED RECORD:", name)
		check(err)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
