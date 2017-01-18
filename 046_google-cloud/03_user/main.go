package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/admin/", admin)
}

func index(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	url, err := user.LogoutURL(ctx, "/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}

func admin(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	url, _ := user.LogoutURL(ctx, "/")
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, `Welcome ADMIN, %s! (<a href="%s">sign out</a>)`, u, url)
}
