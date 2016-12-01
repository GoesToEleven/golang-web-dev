In func "bar" we have code which restricts access to "bar" to only those who have a session:


```
func bar(w http.ResponseWriter, req *http.Request) {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

```
	
You can only access bar if you have a cookie, and if that cookie is associated with a session.

Refactor this code so that the above logic is in a func with the following identifier and signature:

```
func protected(req *http.Request) (user, error) 
```

Make sure func protected works in func bar.
