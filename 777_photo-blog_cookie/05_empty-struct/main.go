package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	data := foo()
	code := getCode(data)
	cookie, err := req.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String() + "|" + data + "|" + code,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Fprintln(res, cookie.Value)
	xs := strings.Split(cookie.Value, "|")
	for i, v := range xs {
		fmt.Fprintln(res, strconv.Itoa(i)+" - "+v)
	}
}

func foo() string {
	m := model{}

	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return string(bs)
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
