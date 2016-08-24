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

func index(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	item1 := memcache.Item{
		Key:        "foo",
		Value:      []byte("bar"),
		Expiration: 10 * time.Second,
	}

	memcache.Set(ctx, &item1)

	item, _ := memcache.Get(ctx, "foo")
	if item != nil {
		fmt.Fprintln(res, string(item.Value))
	}
}
