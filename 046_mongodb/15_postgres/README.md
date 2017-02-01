# THIS IS POSTGRES

We are going to rebuild our postgres example using mongo instead.

First we need to extract our the records from the books table in our bookstore database.

The code in **this folder** has a new route which will give you the following json:

```
[{"Isbn":"978-1505255607","Title":"The Time Machine","Author":"H. G. Wells","Price":5.99},{"Isbn":"978-1503261960","Title":"Wind Sand \u0026 Stars","Author":"Antoine","Price":14.99},{"Isbn":"978-1503261961","Title":"West With The Night","Author":"Beryl Markham","Price":14.99}]
```

notice \u0026

That is valid JSON encoding. That is the unicode escape. & doesn't have to be encoded, but Go does this so we can just let it be. JSON encoders / decoders know how to deal with it.

# To see this JSON output

### turn on postgres
Make sure postgres is on.

### run your server
```
go run *.go
```

### get the json of the books
```
curl -i localhost:8080/books/json
```