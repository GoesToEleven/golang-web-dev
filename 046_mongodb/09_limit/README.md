# limit

### setup

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

### limit
```
db.<collection name>.find(<selection criteria>).limit(n)
```

```
db.crayons.find().limit(3)
```

```
db.customers.find({age:{$gt:32}},{_id:0,name:1,age:1}).limit(2)
```