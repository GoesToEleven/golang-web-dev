# collection commands

### create implicitly
```
db.<collection name>.insert({"name":"McLeod"})
```

### create explicitly
```
db.createCollection(<name>, {<optional options>})
```

#### optional options
| option | type | description |
| --- | --- | --- |
| capped | bool | caps the size |
| size | number | sets size of cap in bytes |
| max | bool | maximum number of documents allowed in capped collection |

[other options including validation](https://docs.mongodb.com/manual/reference/method/db.createCollection/)

#### examples
```
db.createCollection("customers")
```

```
db.createCollection("crs",{capped:true, size:65536,max:1000000})
```

### view collections
```
show collections
```

### drop
```
db.<collection name>.drop()
```