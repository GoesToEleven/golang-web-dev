package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/user"
)

func init() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ctx := appengine.NewContext(r)
	u := user.Current(ctx)

	globalCount, err := memcache.Increment(ctx, "GLOBAL", 1, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userCount, err := memcache.Increment(ctx, u.Email+".COUNTER", 1, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Global %v - User %v", globalCount, userCount)
}
