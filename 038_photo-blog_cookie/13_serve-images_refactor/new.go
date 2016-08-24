package main

import (
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func newVisitor() *http.Cookie {
	data := initialModel()
	code := getCode(data)
	id, _ := uuid.NewV4()
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id.String() + "|" + data + "|" + code,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie
}

func initialModel() string {
	m := model{
		State: false,
		Pictures: []string{
			"one.jpg",
		},
	}

	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return string(bs)
}
