# Mongo Commands

## 02_db

### create
```
use <db name>
```

### use 
```
use <db name>
```

### see current db
```
db
```

### see all db
```
show dbs
```

### insert document
```
db.<collection name>.insert({"name":"McLeod"})
```
If the collection doesn't exist, it is created

### view collections
```
show collections
```

### view documents
```
db.<collection name>.find()
```

### drop db
```
db.dropDatabase()
```

## 03_collection

### create implicitly
```
db.<collection name>.insert({"name":"McLeod"})
```

### create explicitly
```
db.createCollection(<name>, <optional options>)
```

#### optional options
| option | type | description |
| --- | --- | --- |
| autoindex | bool | creates _id index automatically |
| capped | bool | caps the size |
| size | number | sets size of cap in bytes |
| max | bool | maximum number of documents allowed in capped collection |

#### examples
```
db.createCollection("customers")
```

```
db.createCollection("customers",{autoindex:true})
```

```
db.createCollection("customers",{autoindex:true, capped:true,size:65536,max:1000000})
```

### view collections
```
show collections
```

### drop
```
db.<collection name>.drop()

```

## 04_document

### insert
```
db.<collection name>.insert({document})
```

### insert multiple
```
db.<collection name>.insert(< [{document}, {document}, ..., {document}] >)
```
pass in an array of documents

## 05_query

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

### operators

| operator | syntax | example |
| --- | --- | --- |
| equality | {key:value} | db.customers.find({"name":"McLeod"}).pretty() | 
| less than | {key:{$lt:value}} | db.customers.find({"age":{$lt:20}}).pretty() | 
| less than equals | {key:{$lte:value}} | db.customers.find({"age":{$lte:20}}).pretty() | 
| greater than | {key:{$gt:value}} | db.customers.find({"age":{$gt:20}}).pretty() | 
| greater than equals | {key:{$gte:value}} | db.customers.find({"age":{$gte:20}}).pretty() | 
| not equals | {key:{$ne:value}} | db.customers.find({"age":{$ne:20}}).pretty() | 

### and
```
db.customers.find({$and: [{"name":"McLeod"}, {"age":20}]}).pretty()
db.customers.find({$and: [{"name":"McLeod"}, {"age":{$lt:20}}]}).pretty()
db.customers.find({$and: [{"name":"McLeod"}, {"age":{$gt:20}}]}).pretty()
```

### or
```
db.customers.find({$or: [{"name":"McLeod"}, {"age":20}]}).pretty()
db.customers.find({$or: [{"name":"McLeod"}, {"age":{$lt:20}}]}).pretty()
db.customers.find({$or: [{"name":"McLeod"}, {"age":{$gt:20}}]}).pretty()
```

### and or
```
db.customers.find({"role":"owner", $or: [{"name":"McLeod"}, {"age":20}]}).pretty()
db.customers.find({"role":"owner", $or: [{"name":"McLeod"}, {"age":{$lt:20}}]}).pretty()
db.customers.find({"role":"owner", $or: [{"name":"McLeod"}, {"age":{$gt:20}}]}).pretty()
```
WHERE role = 'owner' AND (name = 'mcleod' OR age = 20)

