# db commands

## create

```
use <db name>
```
The above will use the database of the provided name if it exists, and create it if it doesn't

## use 

```
use <db name>
```

#### example
```
use temp
```

The above will use the database of the provided name if it exists, and create it if it doesn't

## see current db

```
db
```

## see all db

```
show dbs
```
You need to have at least one document in a db for it to be seen.

### insert document

```
db.<collection name>.insert({"name":"McLeod"})
```

example
```
db.dogs.insert({"name":"toby"})
```

### view documents
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

### view collections
```
show collections
```

## drop db
```
db.dropDatabase()
```