package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type person struct {
	First string
	Last  string
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "James",
		Last:  "Bond",
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
	}

	xp := []person{p1, p2}

	err := json.NewEncoder(w).Encode(xp)
	if err != nil {
		fmt.Println(err)
	}
}
