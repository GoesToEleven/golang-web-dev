# 02_db

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

### drop db
```
db.dropDatabase()
```

# 03_collection

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

# 04_document

### insert
```
db.<collection name>.insert({document})
```

### insert multiple
```
db.<collection name>.insert(< [{document}, {document}, ..., {document}] >)
```
pass in an array of documents

# 05_query

