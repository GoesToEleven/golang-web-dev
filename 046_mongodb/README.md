# Mongo Commands

## 02_db

#### create

```
use <db name>
```
The above will use the database of the provided name if it exists, and create it if it doesn't

#### use

```
use <db name>
```

#### example
```
use temp
```

The above will use the database of the provided name if it exists, and create it if it doesn't

#### see current db

```
db
```

#### see all db

```
show dbs
```
You need to have at least one document in a db for it to be seen.

#### insert document

```
db.<collection name>.insert({"name":"McLeod"})
```

example
```
db.dogs.insert({"name":"toby"})
```

#### view documents
```
db.<collection name>.find()
```

example
```
db.cats.insert({"firstname":"coco"})
```

```
db.cats.find().pretty()
```

#### view collections
```
show collections
```

#### drop db
```
db.dropDatabase()
```

## 03_collection

# collection commands

#### create implicitly
```
db.<collection name>.insert({"name":"McLeod"})
```

#### create explicitly
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

#### view collections
```
show collections
```

#### drop
```
db.<collection name>.drop()
```

## 04_document

#### insert
```
db.<collection name>.insert({document})
```

#### insert multiple
```
db.<collection name>.insert(< [{document}, {document}, ..., {document}] >)
```
pass in an array of documents

#### example
```
use playroom
```

```
db
```

```
show dbs
```

```
db.crayons.insert([
                    {
                        "hex": "#EFDECD",
                        "name": "Almond",
                        "rgb": "(239, 222, 205)"
                    },
                    {
                        "hex": "#CD9575",
                        "name": "Antique Brass",
                        "rgb": "(205, 149, 117)"
                    },
                    {
                        "hex": "#FDD9B5",
                        "name": "Apricot",
                        "rgb": "(253, 217, 181)"
                    },
                    {
                        "hex": "#78DBE2",
                        "name": "Aquamarine",
                        "rgb": "(120, 219, 226)"
                    },
                    {
                        "hex": "#87A96B",
                        "name": "Asparagus",
                        "rgb": "(135, 169, 107)"
                    },
                    {
                        "hex": "#FFA474",
                        "name": "Atomic Tangerine",
                        "rgb": "(255, 164, 116)"
                    },
                    {
                        "hex": "#FAE7B5",
                        "name": "Banana Mania",
                        "rgb": "(250, 231, 181)"
                    }
                ])
```
[source of crayon json](https://gist.githubusercontent.com/jjdelc/1868136/raw/c9160b1e60bd8c10c03dbd1a61b704a8e977c46b/crayola.json)

```
show collections
```

```
db.crayons.find()
```

```
db.crayons.drop()
```

```
db.dropDatabase()
```

## 05_query
