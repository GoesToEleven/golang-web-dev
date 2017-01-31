# collection commands

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

### drop
```
db.<collection name>.drop()

```