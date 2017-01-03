# Don't run this code yet

We create an ObjectId using the bson package. 

We do this in controllers/user.go in func CreateUser

```
	// create bson ID
	u.Id = bson.NewObjectId()

```

Second, we store the user in mongodb.

We do this in controllers/user.go in func CreateUser

```
uc.session.DB("go-web-dev-db").C("users").Insert(u)
```