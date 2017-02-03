# locking down your database

## create admin super user

```
use admin
db.createUser(
  {
    user: "jamesbond",
    pwd: "moneypennyrocks007sworld",
    roles: [ { role: "userAdminAnyDatabase", db: "admin" } ]
  }
)
```

[built in user roles](https://docs.mongodb.com/manual/reference/built-in-roles/)

#### exit mongo & then start again
```
mongod --auth
```

```
mongo -u "jamesbond" -p "moneypennyrocks007sworld" --authenticationDatabase "admin"
```

#### see current user

```
db.runCommand({connectionStatus : 1})
```

## create regular user
Give this user readwrite permissions on the ```store``` db.

```
db.createUser(
  {
    user: "bond",
    pwd: "moneypenny007",
    roles: [ { role: "readWrite", db: "store" } ]
  }
)
```

#### exit mongo & then start again
```
mongod --auth
```

```
mongo -u "bond" -p "moneypenny007" --authenticationDatabase "store"
```

#### see current user

```
db.runCommand({connectionStatus : 1})
```

#### lock down the database

[enable auth](https://docs.mongodb.com/master/tutorial/enable-authentication/)

[getting auth running on mongo](https://docs.mongodb.com/manual/tutorial/enable-authentication/)

#### exit mongo & then start again with auth enabled
```
mongod --auth
```

```
mongo -u "bond" -p "moneypenny007" --authenticationDatabase "store"
```

#### test

```
use store
```

```
show collections
```

```
db.customers.find()
```

```
db.customers.insert({"role" : "double-zero", "name" : "Elon Musk", "age" : 47 })
```

#### test

launch a new terminal window

```
mongo
```

should be unauthorized:
```
show collections
```

#### drop user
```
db.dropUser("<user name>")
```