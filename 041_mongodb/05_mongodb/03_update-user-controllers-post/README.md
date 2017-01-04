# Don't run this code yet

We create an ObjectId using the bson package. 

IMPORTANT:
Make sure you update your import statements to import packages from the correct location!

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