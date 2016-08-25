package main

import (
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func newVisitor() *http.Cookie {
	mm := initialModel()
	id, _ := uuid.NewV4()
	return makeCookie(mm, id.String())
}

func currentVisitor(m model, id string) *http.Cookie {
	mm := marshalModel(m)
	return makeCookie(mm, id)
}

func makeCookie(mm, id string) *http.Cookie {
	code := getCode(mm)
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id + "|" + mm + "|" + code,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie
}

func marshalModel(m model) string {
	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return string(bs)
}

func initialModel() string {
	m := model{
		State: false,
		Pictures: []string{
			"one.jpg",
		},
	}
	return marshalModel(m)
}
