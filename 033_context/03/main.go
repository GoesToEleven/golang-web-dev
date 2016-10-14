package main

import (
	"net/http"
	"fmt"
	"log"
	"context"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (int, error) {

	ctx, cancel := context.WithTimeout(ctx, 1 * time.Second)
	defer cancel()

	uid := ctx.Value("userID").(int)

	// is this taking too long
	time.Sleep(2 * time.Second)
	if err := ctx.Err(); err != nil {
		return 0, err
	}


	return uid, nil
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

// per request variables
// good candidate for putting into context