package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
	}

	xp := []person{p1, p2}

	fmt.Printf("go data: %+v\n", xp)

	bs, err := json.Marshal(xp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json:", string(bs))
}
