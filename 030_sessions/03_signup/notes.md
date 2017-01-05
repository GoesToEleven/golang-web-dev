# Step 1

Created ```func getUser``` and put it in a new file, session.go. This refactor allows us to use the same code in index and bar.

```
func getUser(w http.ResponseWriter, req *http.Request) user 
```

  
 ## IMPORTANT
 Now that you have two go files, you cannot use "go run main.go" to get your application to run. That is only asking for one go file: main.go. You need to use either "go build" and then run the executable, or "go run *.go"
 
 ## WebStorm Users
 Note to WebStorm users: when you create a new go page that has code in package main, webstorm will highlight an error "multiplate packages in directory"; this will go away in time, or you can restart webstorm for it to go away immediately.
 
# Step 2

Created ```func signup``` and removed the signup code from ```func index```. A new field for password was added to the user struct.

```
func signup(w http.ResponseWriter, req *http.Request)
```

# Step 3

Created ```func alreadyLoggedIn``` and put it on the session.go page. This refactor allows us to use the same code in bar and signup.

```
func alreadyLoggedIn(req *http.Request) bool
```

