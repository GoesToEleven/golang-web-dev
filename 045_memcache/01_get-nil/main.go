package main

import (
	"fmt"
	"net/http"
)
import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

func init() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	item, err := memcache.Get(ctx, "some-key")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, item)
}
