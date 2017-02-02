# Mongo Commands

## 02_db

#### create

```
use <db name>
```
The above will use the database of the provided name if it exists, and create it if it doesn't

#### use

```
use <db name>
```

#### example
```
use temp
```

The above will use the database of the provided name if it exists, and create it if it doesn't

#### see current db

```
db
```

#### see all db

```
show dbs
```
You need to have at least one document in a db for it to be seen.

#### insert document

```
db.<collection name>.insert({"name":"McLeod"})
```

example
```
db.dogs.insert({"name":"toby"})
```

#### view documents
```
db.<collection name>.find()
```

example
```
db.cats.insert({"firstname":"coco"})
```

```
db.cats.find().pretty()
```

#### view collections
```
show collections
```

#### drop db
```
db.dropDatabase()
```

## 03_collection

# collection commands

#### create implicitly
```
db.<collection name>.insert({"name":"McLeod"})
```

#### create explicitly
```
db.createCollection(<name>, {<optional options>})
```

#### optional options
| option | type | description |
| --- | --- | --- |
| capped | bool | caps the size |
| size | number | sets size of cap in bytes |
| max | bool | maximum number of documents allowed in capped collection |

[other options including validation](https://docs.mongodb.com/manual/reference/method/db.createCollection/)

#### examples
```
db.createCollection("customers")
```

```
db.createCollection("crs",{capped:true, size:65536,max:1000000})
```

#### view collections
```
show collections
```

#### drop
```
db.<collection name>.drop()
```

## 04_document


## 05_query
