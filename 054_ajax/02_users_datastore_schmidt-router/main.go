package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
)

type Profile struct {
	Email     string
	FirstName string
	LastName  string
	Age       int
}

var tpl *template.Template

func init() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/api/profile", getAPIProfile)
	router.POST("/api/profile", updateAPIProfile)
	http.Handle("/", router)

	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func index(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "index.html", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}

func getAPIProfile(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	key := datastore.NewKey(ctx, "Profile", u.Email, 0, nil)
	var profile Profile
	err := datastore.Get(ctx, key, &profile)
	if err != nil {
		profile.Email = u.Email
	}
	json.NewEncoder(res).Encode(profile)
}

func updateAPIProfile(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var profile Profile
	json.NewDecoder(req.Body).Decode(&profile)

	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	key := datastore.NewKey(ctx, "Profile", u.Email, 0, nil)
	_, err := datastore.Put(ctx, key, &profile)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}
