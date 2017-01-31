# query

### setup

```
use store
db
show dbs
db.customers.insert([{"role":"double-zero","name": "Bond","age": 32},{"role":"citizen","name": "Moneypenny","age":24},{"role":"citizen","name": "Q","age":67},{"role":"citizen","name": "M","age":57},{"role":"citizen","name": "Dr. No","age":52}])
```

### find
```
db.<collection name>.find()
db.customers.find()
```

### find one
```
db.<collection name>.findOne()
db.customers.findOne()
```

### find specific
```
db.customers.find({name:"Bond"})
db.customers.find({name:"Bond"})
```
You can use "name" or name.
JSON: specification is to enclose name (object name-value pair) in double qoutes
[JSON specification](json-spec.html)

### and
```
db.customers.find({$and: [{name:"Bond"}, {age:32}]})
db.customers.find({$and: [{name:"Bond"}, {age:32}]}).pretty()
db.customers.find({$and: [{name:"Bond"}, {age:{$lt:20}}]})
db.customers.find({$and: [{name:"Bond"}, {age:{$gt:20}}]})
```

### or
```
db.customers.find({$or: [{name:"Bond"}, {age:24}]})
db.customers.find({$or: [{name:"Bond"}, {age:{$lt:20}}]})
db.customers.find({$or: [{name:"Bond"}, {age:{$gt:20}}]})
```

### and or
```
db.customers.find({"role":"double-zero"})
db.customers.find({age:{$gt:32}})
db.customers.find({$and: [$or:[{"role":"double-zero"}], $or: [{"age":{$gt:32}}]]})

db.customers.find( {
    $and : [
        { $or : [ { role : 0.99 }, { price : 1.99 } ] },
        { $or : [ { sale : true }, { qty : { $lt : 20 } } ] }
    ]
} )



db.customers.find({"role":"double-zero", $or: [{"age":{$gt:32}}]})
db.customers.find({"role":"citizen", $or: [{"name":"M"}, {"age":{$lt:20}}]})
db.customers.find({"role":"citizen", $or: [{"name":"M"}, {"age":{$gt:20}}]})
```
WHERE role = 'owner' AND (name = 'mcleod' OR age = 20)



### pretty
```
db.<collection name>.find().pretty()
```
pretty prints the results

### operators

| operator | syntax | example |
| --- | --- | --- |
| equality | {key:value} | db.customers.find({"name":"Bond"}).pretty() | 
| less than | {key:{$lt:value}} | db.customers.find({"age":{$lt:20}}).pretty() | 
| less than equals | {key:{$lte:value}} | db.customers.find({"age":{$lte:20}}).pretty() | 
| greater than | {key:{$gt:value}} | db.customers.find({"age":{$gt:20}}).pretty() | 
| greater than equals | {key:{$gte:value}} | db.customers.find({"age":{$gte:20}}).pretty() | 
| not equals | {key:{$ne:value}} | db.customers.find({"age":{$ne:20}}).pretty() | 