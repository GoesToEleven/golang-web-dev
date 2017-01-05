package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var a []string

	rcvd := `null`
	err := json.Unmarshal([]byte(rcvd), &a)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
