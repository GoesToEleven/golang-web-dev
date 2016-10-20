package main

type client struct {
	Email string
	Fname string
	Loggedin bool
}

var db = make(map[string]client)
