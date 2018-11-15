package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Book struct {
	ISBN string
	Title string
	Author string
	Price float32
}

func main() {
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

	rs, err := db.Query("select * from books;")

	xb := make([]Book, 0)
	for rs.Next() {
		b := Book{}
		rs.Scan(&b.ISBN, &b.Title, &b.Author, &b.Price)
		xb = append(xb, b)

		if err = rs.Err(); err != nil {
				panic(err)
		}
	}

	for i, v := range xb {
		fmt.Println(i, v)
	}
}