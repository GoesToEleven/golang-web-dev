# create index

```
db.<collection name>.createIndex({<field to index>:<1 for ascend, -1 descend>})
```

create index
```
db.oscars.createIndex({title:1})
```

see indexes
```
db.oscars.getIndexes()
```

[learn to create a unique index and more](https://docs.mongodb.com/manual/reference/method/db.collection.createIndex/#db.collection.createIndex)