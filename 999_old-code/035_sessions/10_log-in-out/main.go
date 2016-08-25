package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("logged-in")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "logged-in",
			Value: "0",
			//Secure: true,
			HttpOnly: true,
		}
	}

	// check log in
	if req.Method == "POST" {
		password := req.FormValue("password")
		if password == "secret" {
			cookie = &http.Cookie{
				Name:  "logged-in",
				Value: "1",
				//Secure: true,
				HttpOnly: true,
			}
		}
	}

	// if logout, then logout
	if req.URL.Path == "/logout" {
		cookie = &http.Cookie{
			Name:   "logged-in",
			Value:  "0",
			MaxAge: -1,
			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
		http.Redirect(res, req, "/", 303)
		return
	}

	http.SetCookie(res, cookie)
	var html string

	// not logged in
	if cookie.Value == "0" {
		html = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
			<h1>LOG IN</h1>
			<form method="POST">
				<h3>Password</h3>
				<input type="text" name="password">
				<br>
				<input type="submit">
			</form>
			</body>
			</html>`
	}

	// logged in
	if cookie.Value == "1" {
		html = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
			<h1><a href="/logout">LOG OUT</a></h1>
			</body>
			</html>`
	}

	io.WriteString(res, html)
}

// NOT GOOD PRACTICE
// adding user data to a cookie
// with no way of knowing whether or not
// they might have altered that data
//
// HMAC would allow us to determine
// whether or not the data in the cookie was altered
//
// however, best to store user data
// on the server
// and keep backups
