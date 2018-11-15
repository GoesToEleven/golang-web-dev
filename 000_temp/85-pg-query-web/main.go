package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

var tpl *template.Template
var db *sql.DB

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	connStrn := "host=localhost port=5432 dbname=bookstore user=bond password=password sslmode=disable"

	db, err := sql.Open("postgres", connStrn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
}

type Book struct {
	Isbn string
	Title string
	Author string
	Price float32
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	rs, err := db.Query("select * from books;")

	xb := make([]Book, 0)
	for rs.Next() {
		b := Book{}
		rs.Scan(&b.Isbn, &b.Title, &b.Author, &b.Price)
		xb = append(xb, b)

		if err = rs.Err(); err != nil {
			panic(err)
		}
	}
	tpl.ExecuteTemplate(w, "books.gohtml", xb)
}



//rs, err := db.Query("select * from books;")
//if err != nil {
//	http.Error(w, http.StatusText(500), 500)
//	return
//}
//defer rs.Close()
//
//xb := make([]Book, 0)
//for rs.Next() {
//	b := Book{}
//	err := rs.Scan(&b.ISBN, &b.Title, &b.Author, &b.Price)
//	if err != nil {
//		http.Error(w, http.StatusText(500), 500)
//		return
//
//	}
//	xb = append(xb, b)
//
//	if err = rs.Err(); err != nil {
//		http.Error(w, http.StatusText(500), 500)
//		return
//	}
//}
//
//tpl.ExecuteTemplate(w, "index.gohtml", xb)
