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
	j := `[{"First":"James","Last":"Bond"},{"First":"Miss","Last":"Moneypenny"}]`
	fmt.Println("json:", j)

	xp := []person{}

	err := json.Unmarshal([]byte(j), &xp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("go data: %+v\n", xp)

	for i, v := range xp {
		fmt.Println(i, v)
		fmt.Println("\t", v.First)
	}
}
