package main

import (
	"net/http"
	"log"
)

func about(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}

