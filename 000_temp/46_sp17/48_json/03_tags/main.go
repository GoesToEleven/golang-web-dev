package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"fname"`
	Last  string
}

func main() {
	j := `[{"fname":"James","last":"Bond"},{"fname":"Miss","last":"Moneypenny"}]`
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
