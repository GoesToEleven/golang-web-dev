package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type person struct {
	First string
	Last  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/send", send)
	http.HandleFunc("/catch", catch)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	xp := getUsers(r)

	tpl.ExecuteTemplate(w, "index.gohtml", xp)
}

func send(w http.ResponseWriter, r *http.Request) {
	xp := getUsers(r)

	tpl.ExecuteTemplate(w, "send.gohtml", xp)
}

func catch(w http.ResponseWriter, r *http.Request) {
	xp := getUsers(r)

	if r.Method == http.MethodPost {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal(bs, &xp)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(xp)

		s := base64.StdEncoding.EncodeToString(bs)
		fmt.Println(s)

		c := &http.Cookie{
			Name:     "userdata",
			Value:    s,
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(w, c)
	}

	if r.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "catch.gohtml", xp)
	}
}

func getUsers(r *http.Request) []person {
	var xp []person

	c, err := r.Cookie("userdata")
	if err == nil {
		bs, err := base64.StdEncoding.DecodeString(c.Value)
		if err != nil {
			fmt.Println(err)
			return xp
		}

		fmt.Println("just after base64 decode:", string(bs))

		err = json.Unmarshal(bs, &xp)
		if err != nil {
			fmt.Println("unmarshalling in decode from b64", err)
		}
	}
	return xp
}
