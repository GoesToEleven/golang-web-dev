# STEP 1:

There is a page "session.go"

The "session.go" page has a func with the following identifier and signature:

```
getSession(w http.ResponseWriter, req *http.Request) user 
```

Refactor your code on "main.go" to use func getSession where appropriate.

# STEP 2:

Put func protected onto this "session.go" page also.


# STEP 3: 

The "session.go" page also has a func with the following identifier and signature:

```
func alreadyLoggedIn(req *http.Request) bool
```

Refactor your code on "main.go" to use func alreadyLoggedIn where appropriate.

## IMPORTANT
Now that you have two go files, you cannot use "go run main.go" to get your application to run. That is only asking for one go file: main.go. You need to use either "go build" and then run the executable, or "go run *.go"

## WebStorm Users
Note to WebStorm users: when you create a new go page that has code in package main, webstorm will highlight an error "multiplate packages in directory"; this will go away in time, or you can restart webstorm for it to go away immediately.