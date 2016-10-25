package main

var db map[string]*User

func init() {
	db = make(map[string]*User)
}
