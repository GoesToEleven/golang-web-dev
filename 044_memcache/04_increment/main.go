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

func index(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	ctx := appengine.NewContext(req)
	u := user.Current(ctx)

	globalCount, err := memcache.Increment(ctx, "GLOBAL", 1, 0)
	if err != nil {
		fmt.Println("1", err)
	}
	userCount, err := memcache.Increment(ctx, u.Email+".COUNTER", 1, 0)
	if err != nil {
		fmt.Println("2", err)
	}

	fmt.Fprintln(res, "Global", globalCount)
	fmt.Fprintln(res, "User", userCount)
}
