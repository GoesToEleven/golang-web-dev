package main

import (
	"net/http"
	"html/template"
	"strconv"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := cookieCounter(r)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	http.SetCookie(w, c)

	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}


func bar(w http.ResponseWriter, r *http.Request) {
	c, err := cookieCounter(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	http.SetCookie(w, c)

	tpl.ExecuteTemplate(w, "about.gohtml", c.Value)
}

func cookieCounter(r *http.Request) (*http.Cookie, error) {
	var c *http.Cookie

	c, err := r.Cookie("wackadoodle")
	if err != nil {
		c = &http.Cookie{
			Name: "wackadoodle",
			Value: "0",
			Path: "/",
		}
	}

	i, err := strconv.Atoi(c.Value)
	if err != nil {
		return c, err
	}
	i++
	c.Value = strconv.Itoa(i)

	return c, nil
}
