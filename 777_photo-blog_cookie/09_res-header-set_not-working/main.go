package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"io"
	"net/http"
	"strings"
)

type model struct {
	State    bool
	Pictures []string
}

var tpl *template.Template

func init() {
	tpl, _ = template.ParseGlob("templates/*.html")
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
	// usrToken := xs[0]
	usrPics := xs[1]
	usrCode := xs[2]
	fmt.Fprintln(res, data)
	fmt.Fprintln(res, usrPics)
	if usrCode == getCode(usrPics) {
		fmt.Fprintln(res, "Code valid")
	} else {
		fmt.Fprintln(res, "Code Invalid")
	}

	var m model
	err = json.Unmarshal([]byte(usrPics), &m)
	if err != nil {
		fmt.Println("error unmarshalling: ", err)
	}

	fmt.Fprintln(res, m)
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(res, "index.html", m)
}

func foo() string {
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

	return string(bs)
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
