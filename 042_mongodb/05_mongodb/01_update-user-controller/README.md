# Install Mongo

# Go get driver for mongo

```
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson
```

# In this step:

Don't run this code

Just making updates - a several step process.

We will need a mongo session to use in the CRUD methods.

We need our UserController to have access to a mongo session.

Let's add this to controllers/user.go

```
UserController struct {  
    session *mgo.Session
}
```

And now add this to controllers/user.go

```
func NewUserController(s *mgo.Session) *UserController {  
    return &UserController{s}
}
```

And now add this to main.go

```
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
```

and this

```
uc := controllers.NewUserController(getSession())  
```

1. Enter this at the terminal

```
curl http://localhost:8080/user/1
```