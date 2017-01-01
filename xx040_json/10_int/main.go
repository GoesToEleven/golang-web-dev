package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var data int
	rcvd := `42`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)
}
