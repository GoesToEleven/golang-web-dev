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

### view collections
```
db
```

```
show collections
```

## drop db
```
db.dropDatabase()
```