# Note

**Remember:** You need to capitalize to have a variable, type, or func exported from a package

# run the application and make a request
```
curl -i localhost:8080/books
```

```
curl -i -X POST -d "isbn=978-1470184841&title=Metamorphosis&author=Franz Kafka&price=5.90" localhost:8080/books/create/process
```