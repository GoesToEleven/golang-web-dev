package main

import (
	"net/http"
	"log"
)

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
