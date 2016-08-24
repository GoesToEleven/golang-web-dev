# Logic

### Everyone Has A Cookie & Session ID

With every request, we ask for the cookie. In the cookie, a session ID (uuid) is stored.

```go
cookie := getCookie(res, req)
```

```go
func getCookie(res http.ResponseWriter, req *http.Request) *http.Cookie {

	cookie, err := req.Cookie("session-id")

	if err != nil {
		cookie = newVisitor(req)
		http.SetCookie(res, cookie)
		return cookie
	}

	return cookie
}
```

```go
func newVisitor(req *http.Request) *http.Cookie {
	m := initialModel()
	id, _ := uuid.NewV4()
	return makeCookie(m, id.String(), req)
}

func initialModel() model {
	m := model{
		Name:  "",
		State: false,
		Pictures: []string{
			"one.jpg",
		},
	}
	return m
}
```

```go
func makeCookie(m model, id string, req *http.Request) *http.Cookie {

	// DATASTORE
	storeDstore(m, id, req)

	// MEMCACHE
	storeMemc(m, id, req)

	// COOKIE
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie
}
```

Notice that every time a cookie is made by `func makeCookie`, we are also storing data in the datastore and in memcache.

This is important logic.

If we ever ask for data from the datastore and are unable to locate it for a user, then that user needs to begin at the beginning of our process again and receive a new cookie.

Whenever we issue a user ID (uuid), that ID will also have corresponding data in the datastore.



If there is no cookie, we make a cookie and set it.

### Error Handling Helps Logic Flow

Notice the additions of error handling and how this helps the flow of the program's logic.