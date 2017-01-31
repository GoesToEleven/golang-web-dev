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