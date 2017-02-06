# update & save

update will update a record

save will overwrite a record

### update
```
db.<collection name>.update(<selection criteria>, <update data>, <optional options>)
```

example
```
db.customers.find()
```

gives this data
```
{ "_id" : ObjectId("5891221756867ebff44cc885"), "role" : "double-zero", "name" : "Bond", "age" : 32 }
{ "_id" : ObjectId("5891221756867ebff44cc886"), "role" : "citizen", "name" : "Moneypenny", "age" : 32 }
{ "_id" : ObjectId("5891221756867ebff44cc887"), "role" : "citizen", "name" : "Q", "age" : 67 }
{ "_id" : ObjectId("5891221756867ebff44cc888"), "role" : "citizen", "name" : "M", "age" : 57 }
{ "_id" : ObjectId("5891221756867ebff44cc889"), "role" : "citizen", "name" : "Dr. No", "age" : 52 }
```

update like this
```
db.customers.update({_id:ObjectId("5891221756867ebff44cc886")},{$set:{role:"double-zero"}})
```

```
db.customers.update({name:"Moneypenny"},{$set:{role:"double-zero"}})
```

```
db.customers.update({name:"Moneypenny"},{$set:{role:"citizen", name: "Miss Moneypenny"}})
```

```
db.customers.update({age:{$gt:35}},{$set:{role:"double-zero"}})
```

```
db.customers.update({age:{$gt:35}},{$set:{role:"double-zero"}}, {multi:true})
```
[see options](https://docs.mongodb.com/manual/reference/method/db.collection.update/)

```
db.customers.update({},{$set:{role:"citizen"}}, {multi:true})
```
[see query documentation](https://docs.mongodb.com/manual/tutorial/query-documents/)

("5893888012acb8ada532a8e4"

### save
```
db.customers.save({"role":"villain","name":"Jaws","age":43})
```

```
db.customers.save({"_id":ObjectId("5891221756867ebff44cc889"),"role":"villain","name":"Goldfinger","age":77})
```

```
db.customers.save({"_id":ObjectId("5893888012acb8ada532a8e4"),"role":"villain","name":"PussyGalore","age":31})
```