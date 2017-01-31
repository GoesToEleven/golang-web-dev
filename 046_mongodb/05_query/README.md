# query

### find
```
db.<collection name>.find()
```

### find one
```
db.<collection name>.findOne()
```

### pretty
```
db.<collection name>.find().pretty()
```
pretty prints the results

### insert multiple
```
db.<collection name>.insert(< [{document}, {document}, ..., {document}] >)
```
pass in an array of documents

### operators

| operator | syntax | example |
| --- | --- | --- |
| equality | {<key>:<value>} | db.customers.find({"name":"McLeod"}).pretty() | 
| less than | {<key>:{$lt:<value>}} | db.customers.find({"score":{$lt:20}}).pretty() | 
| less than equals | {<key>:{$lte:<value>}} | db.customers.find({"score":{$lte:20}}).pretty() | 
| greater than | {<key>:{$gt:<value>}} | db.customers.find({"score":{$gt:20}}).pretty() | 
| greater than equals | {<key>:{$gte:<value>}} | db.customers.find({"score":{$gte:20}}).pretty() | 
| not equals | {<key>:{$ne:<value>}} | db.customers.find({"score":{$ne:20}}).pretty() | 

###
```
```

###
```
```

###
```
```

###
```
```

###
```
```

###
```
```