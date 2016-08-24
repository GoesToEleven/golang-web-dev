package main

import (
	"encoding/base64"
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

func makeCookie(mm []byte, id string) *http.Cookie {
	b64 := base64.URLEncoding.EncodeToString(mm)
	code := getCode(b64)
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id + "|" + b64 + "|" + code,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie
}

func marshalModel(m model) []byte {
	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return bs
}

func initialModel() []byte {
	m := model{
		State: false,
		Pictures: []string{
			"one.jpg",
		},
	}
	return marshalModel(m)
}
