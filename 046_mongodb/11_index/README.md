# create index

```
db.<collection name>.createIndex({<field to index>:<1 for ascend, -1 descend>})
```

```
db.oscars.createIndex({title:1})
db.oscars.createIndex({releaseYear:1, releaseDay:1})
```

[learn to create a unique index and more](https://docs.mongodb.com/manual/reference/method/db.collection.createIndex/#db.collection.createIndex)