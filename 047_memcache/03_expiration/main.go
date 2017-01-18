package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"time"
)

func init() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	item1 := memcache.Item{
		Key:        "foo",
		Value:      []byte("bar"),
		Expiration: 10 * time.Second,
	}

	memcache.Set(ctx, &item1)

	item, err := memcache.Get(ctx, "foo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(item.Value))
}
