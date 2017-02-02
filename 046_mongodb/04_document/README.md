# document commands

### insert
```
db.<collection name>.insert({document})
```

### insert multiple
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
