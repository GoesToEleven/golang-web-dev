package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type model struct {
	State    bool
	Pictures []string
}

func Model(s string) model {
	xs := strings.Split(s, "|")
	usrData := xs[1]
	var m model
	err := json.Unmarshal([]byte(usrData), &m)
	if err != nil {
		fmt.Println("error unmarshalling: ", err)
	}
	return m
}
