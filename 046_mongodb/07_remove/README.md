# remove document

```
db.<collection name>.remove(<selection criteria>)
db.customers.remove({role:"double-zero"})
db.customers.remove({role:"villain"})
```
removes all it matches

### remove only 1
```
db.customers.remove({role:"citizen"},1)
```

### remove
```
db.customers.remove({role:"citizen"})
```

### put documents back
```
db.customers.insert([{"role":"double-zero","name": "Bond","age": 32},{"role":"citizen","name": "Moneypenny","age":32},{"role":"citizen","name": "Q","age":67},{"role":"citizen","name": "M","age":57},{"role":"citizen","name": "Dr. No","age":52}])
```

### remove all
```
db.customers.remove({})
```

### put documents back
```
db.customers.insert([{"role":"double-zero","name": "Bond","age": 32},{"role":"citizen","name": "Moneypenny","age":32},{"role":"citizen","name": "Q","age":67},{"role":"citizen","name": "M","age":57},{"role":"citizen","name": "Dr. No","age":52}])
```