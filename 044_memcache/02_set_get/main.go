package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

func init() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	item1 := memcache.Item{
		Key:   "foo",
		Value: []byte("bar"),
	}

	memcache.Set(ctx, &item1)

	item, err := memcache.Get(ctx, "foo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(item.Value))
}
