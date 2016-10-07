package main

import (
	"log"
	"net/http"
)

func HandleError(res http.ResponseWriter, e error) {
	if e != nil {
		http.Error(res, e.Error(), http.StatusInternalServerError)
		log.Println(e)
	}
}
