package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var data []int
	fmt.Println(data)

	rcvd := `null`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}
