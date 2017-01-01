# Step 1: 
Created type session which is a struct with a lastActivity field. This will allow us to know the last time a session was used.

# Step 2:
Updated dbSessions to be of type map[string]session

# Step 3: 
Updated all reads/writes to dbSessions

# Step 4:
Apply the MaxAge field to cookie

# Step 5:
Updated func alreadyLoggedIn to be able to set a cookie, adding the ResponseWriter to its parameters

```
func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
```

# Step 6:
Added dbSessionsCleaned time.Time to keep track of the last time we cleaned out our sessions.

# Step 7:
Added func cleanSessions to remove unused sessions from dbSessions. Set it to run whenever someone logs out and a certain amount of time has elapsed (in production you'd set this to run during a time when the server wasn't busy).


