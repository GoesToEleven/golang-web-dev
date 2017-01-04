Don't run this code

Just making updates - a several step process.

IMPORTANT:
Make sure you update your import statements to import packages from the correct location!

In this step:

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
uc := controllers.NewUserController(getSession())  
```