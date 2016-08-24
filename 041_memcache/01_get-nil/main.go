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

func index(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	item, _ := memcache.Get(ctx, "some-key")
	fmt.Fprintln(res, item)
}
