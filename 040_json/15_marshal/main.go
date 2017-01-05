package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {
	m := model{
		State: true,
		Pictures: []string{
			"one.jpg",
			"two.jpg",
			"three.jpg",
		},
	}

	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	os.Stdout.Write(bs)
}

// Answer
// "... Struct values encode as JSON objects. Each exported struct field becomes a member of the object..."
