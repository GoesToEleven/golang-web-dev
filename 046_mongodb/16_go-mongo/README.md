# setup

```
go get -u gopkg.in/mgo.v2
```

[https://godoc.org/gopkg.in/mgo.v2](https://godoc.org/gopkg.in/mgo.v2)

[http://labix.org/mgo](http://labix.org/mgo)

# setup database

#### restart mongo
close any open mongo connections, then restart with these commands:
```
mongo
```

in a new tab
```
mongod
```

#### create db
```
use bookstore
```

#### create collection books & insert books
```
db.books.insert([{"isbn":"978-1505255607","title":"The Time Machine","author":"H. G. Wells","price":5.99},{"isbn":"978-1503261960","title":"Wind Sand \u0026 Stars","author":"Antoine","price":14.99},{"isbn":"978-1503261961","title":"West With The Night","author":"Beryl Markham","price":14.99}])
```

#### test
```
db.books.find()
```

#### user setup
```
db.createUser(
  {
    user: "bond",
    pwd: "moneypenny007",
    roles: [ { role: "readWrite", db: "bookstore" } ]
  }
)
```

#### exit mongo & then start again with auth enabled
```
mongod --auth
```

```
mongo -u "bond" -p "moneypenny007" --authenticationDatabase "bookstore"
```

#### test

```
use bookstore
```

```
show collections
```

```
db.books.find()
```

```
db.books.insert({"isbn" : "978-1503261777", "title" : "Never Say Never", "author" : "Ian Fleming", "price" : 24.99 })
```

```
db.books.find()
```

# GO & MONGO

#### db access

```
mongodb://myuser:mypass@localhost:27017/dbToAccess
```

If the port number is not provided for a server, it defaults to 27017.

for our example:
```
mongodb://bond:moneypenny007@localhost:27017/bookstore
```
[https://godoc.org/gopkg.in/mgo.v2#Dial](https://godoc.org/gopkg.in/mgo.v2#Dial)

#### db.go
Update to use mongo. You will use the ```mgo.Dial``` to create a session. You can still assign this to the variable ```DB```.



#### models.go
Update to use mongo. 


# run the application and make a request
```
curl -i localhost:8080/books
```

```
curl -i -X POST -d "isbn=978-1470184841&title=Metamorphosis&author=Franz Kafka&price=5.90" localhost:8080/books/create/process
```