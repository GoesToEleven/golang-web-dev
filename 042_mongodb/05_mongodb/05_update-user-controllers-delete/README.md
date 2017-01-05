# In this step:

We will delete a user from mongodb.

This is identical to what we did in the last step to GET a user.

First we will get the user id from the URL

```
id := p.ByName("id")
```

Next we will Verify that the id is an ObjectId

```
if !bson.IsObjectIdHex(id) {
	w.WriteHeader(http.StatusNotFound) // 404
	return
}
```

ObjectIdHex returns an ObjectId from the provided hex representation.

```
	oid := bson.ObjectIdHex(id)
```

Next, add code to delete the user

```
if err := uc.session.DB("go_rest_tutorial").C("users").RemoveId(oid); err != nil {
	w.WriteHeader(404)
	return
}
```

# Run this code

1. Start your server

## DELETE a user from mongodb

Enter this at the terminal

```
curl -X POST -H "Content-Type: application/json" -d '{"name":"Miss Moneypenny","gender":"female","age":27}' http://localhost:8080/user
```

```
curl http://localhost:8080/user/<enter-user-id-here>

```


```
curl -X DELETE http://localhost:8080/user/<enter-user-id-here>
```