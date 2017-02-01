# aggregate

Aggregations operations process data records and return computed results. Aggregation operations group values from multiple documents together, and can perform a variety of operations on the grouped data to return a single result. MongoDB provides three ways to perform aggregation: the aggregation pipeline, the map-reduce function, and single purpose aggregation methods.

[documenation about aggregation](https://docs.mongodb.com/manual/aggregation/)

### single purpose aggregation

[documenation about single purpose aggregation](https://docs.mongodb.com/manual/aggregation/#single-purpose-agg-operations)

There are two functions you can use:

#### [db.collection.count()](https://docs.mongodb.com/manual/reference/method/db.collection.count/#db.collection.count)

syntax
```
db.collection.count(query, options)
```


| Parameter | Description |
| --- | --- | 
| query | The query selection criteria.
| options | Optional. Extra options for modifying the count.


#### [db.collection.distinct()](https://docs.mongodb.com/manual/reference/method/db.collection.distinct/#db.collection.distinct)

syntax
```
db.collection.distinct(field, query, options)
```


| Parameter | Description |
| --- | --- | 
| field | The field for which to return distinct values.
| query | A query that specifies the documents from which to retrieve the distinct values.
| options | Optional. A document that specifies the options. See Options.

### examples - count()
```
db.oscars.count()
```

The above is equivalent to

```
db.oscars.find().count()
```

```
db.customers.find({role:"citizen"}).count()
```

```
db.customers.find({$or: [{name:"Bond"}, {age:{$gt:32}}]}).count()
```

### examples - distinct() - setup
```
db.inventory.insert([
{ "_id": 1, "dept": "A", "item": { "sku": "111", "color": "red" }, "sizes": [ "S", "M" ] },
{ "_id": 2, "dept": "A", "item": { "sku": "111", "color": "blue" }, "sizes": [ "M", "L" ] },
{ "_id": 3, "dept": "B", "item": { "sku": "222", "color": "blue" }, "sizes": "S" },
{ "_id": 4, "dept": "A", "item": { "sku": "333", "color": "black" }, "sizes": [ "S" ] }
])
```

### examples - distinct()

```
db.inventory.distinct( "dept" )
```

embedded field:

```
db.inventory.distinct( "item.sku" )
```

```
db.inventory.distinct( "sizes" )
```


You use different **expressions** (listed below) in aggregate group operations.

![aggregate pipeline](aggregate.png)

```
db.<collection name>.aggregate([{<match, sort, geoNear>},{<group>}])
```

[documenation](https://docs.mongodb.com/manual/aggregation/)
[match, sort, geonear](https://docs.mongodb.com/manual/core/aggregation-pipeline/#aggregation-pipeline-operators-and-performance)

### examples - setup
```
db.orders.insert([
{"cust_id":"A123","amount":500,"status":"A"},
{"cust_id":"A123","amount":250,"status":"A"},
{"cust_id":"B212","amount":200,"status":"A"},
{"cust_id":"A123","amount":300,"status":"D"}
])
```

### examples
```
db.orders.aggregate([
{$match:{status:"A"}},
{$group:{_id: "$cust_id",total: {$sum:"$amount"}}}
])
```


### expressions

| expression | description |

```
db.oscars.createIndex({title:1})
db.oscars.createIndex({releaseYear:1, releaseDay:1})
```

[learn to create a unique index and more](https://docs.mongodb.com/manual/reference/method/db.collection.createIndex/#db.collection.createIndex)