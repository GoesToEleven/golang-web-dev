package dstore

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

type model struct {
	Fname string
}

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/retrieve", noConfusion)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func index(res http.ResponseWriter, req *http.Request) {

	ctx := appengine.NewContext(req)

	id, err := getID(res, req)
	if err != nil {
		log.Errorf(ctx, "ERROR index getID: %s", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// DATASTORE
	m := model{
		Fname: "Todd",
	}
	err = storeDstore(m, id, req)
	if err != nil {
		log.Errorf(ctx, "ERROR index storeDstore: %s", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// MEMCACHE
	err = storeMemc(m, id, req)
	if err != nil {
		log.Errorf(ctx, "ERROR index storeMemc: %s", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, "EVERYTHING SET ID:"+id)
}

func noConfusion(res http.ResponseWriter, req *http.Request) {

	html := `
	<form method="POST">
	    <input type="text" name="koala">
	    <input type="submit" value="submit">
	</form>
	`

	if req.Method == "POST" {
		id := req.FormValue("koala")

		// get value
		m, err := retrieveMemc(id, req)
		if err != nil {
			ctx := appengine.NewContext(req)
			log.Errorf(ctx, "ERROR noConfusion retrieveMemc: %s", err)
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		html += `
			<br>
			<h1>
			Value: ` + m.Fname + `
			</h1>
		`
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(res, html)
}
