package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type model struct {
	State    bool
	Pictures []string
}

func Model(c *http.Cookie) model {
	xs := strings.Split(c.Value, "|")
	usrData := xs[1]
	var m model
	err := json.Unmarshal([]byte(usrData), &m)
	if err != nil {
		fmt.Println("error unmarshalling: ", err)
	}
	return m
}
