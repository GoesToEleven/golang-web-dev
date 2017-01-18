package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
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
