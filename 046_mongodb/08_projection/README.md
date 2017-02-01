# projection
Retrieving part of a document; only some of the fields.

```
db.<collection name>.find(<selection criteria>,<list of fields with toggle 0 or 1>)
```

```
db.customers.find({},{_id:0,name:1,})
```
_id is displayed by default; turn off with 0

```
db.customers.find({},{_id:0,name:1,age:1})
```

```
db.customers.find({age:{$gt:32}},{_id:0,name:1,age:1})
```