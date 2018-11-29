package main

import (
	"html/template"
	"net/http"
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"fmt"
)

type customer struct {
	ID int
	First string
}

var tpl *template.Template
var db *sql.DB

func init() {
	var err error

	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	connStr := "dbname=thanksgiving user=turkey password=gravy host=localhost port=5432 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("**** NO OPEN ****", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("**** NO PING WORKING ****", err)
	} else {
		fmt.Println("Ping successful")
	}
}

func main() {
	defer db.Close()

	http.HandleFunc("/", index)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./resources"))))
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM customers;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	xc := []customer{}

	for rows.Next() {
		c := customer{}
		rows.Scan(&c.ID, &c.First)
		xc = append(xc, c)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", xc)
}
