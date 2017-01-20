# query a single book

add these lines of code:

## func main()
```
  http.HandleFunc("/books/show", booksShow)
```

## func bookShow added
Arguments to the SQL function are referenced in the function body using the syntax $n: $1 refers to the first argument, $2 to the second, and so on. If an argument is of a composite type, then the dot notation, e.g., $1.name, can be used to access attributes of the argument. The arguments can only be used as data values, not as identifiers.[source: postgres docs](https://www.postgresql.org/docs/9.1/static/xfunc-sql.html)

"Behind the scenes, **db.QueryRow (and also db.Query() and db.Exec()) work by creating a new prepared statement** on the database, and subsequently execute that prepared statement using the placeholder parameters provided. This means that **all three methods are safe from SQL injection** when used correctly . From Wikipedia:
 
 Prepared statements are resilient against SQL injection, because parameter values, which are transmitted later using a different protocol, need not be correctly escaped. If the original statement template is not derived from external input, injection cannot occur.
 
The placeholder parameter syntax differs depending on your database. Postgres uses the $N notation, but MySQL, SQL Server and others use the ? character as a placeholder." - Alex Edwards

# run the application and make a request
```
curl -i localhost:8080/books/show?isbn=978-1505255607
```