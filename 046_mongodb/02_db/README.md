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
db.<db name>.insert({"name":"McLeod"})
```

## drop db
```
db.dropDatabase()
```