package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

	connstr := "host=localhost port=5432 dbname=bookstore user=bond password=password sslmode=disable"

	//  "postgres://bond:password@localhost/bookstore?sslmode=disable"

	db, err := sql.Open("postgres", connstr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}
