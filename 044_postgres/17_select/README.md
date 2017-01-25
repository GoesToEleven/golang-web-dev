# our pathway

We will be following the example created in [this article by Alex Edwards](http://www.alexedwards.net/blog/practical-persistence-sql) and licensed under a [MIT license](https://opensource.org/licenses/MIT)

In order to successfully pull records from a table in a database as our user ```bond```, we will need to [ALTER](https://www.postgresql.org/docs/9.6/static/sql-alteruser.html) bond to have a different role.

# alter bond's role
```
alter user bond with superuser;
```

# switch to your bookstore database
You should already have a ```bookstore``` database:

list databases
```
\l
```

switch into that database
```
\c bookstore
```

directory of tables, if any
```
\d
```

# create table
```
CREATE TABLE books (
  isbn    char(14)     PRIMARY KEY NOT NULL,
  title   varchar(255) NOT NULL,
  author  varchar(255) NOT NULL,
  price   decimal(5,2) NOT NULL
);
```

directory of tables
```
\d
```

details of table ```books```
```
\d books
```

# insert records
```
INSERT INTO books (isbn, title, author, price) VALUES
('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);
```

view records
```
SELECT * FROM books;
```

# main.go

## importing the driver
make sure you import the driver
```
_ "github.com/lib/pq"
```
"We don't use anything in the pq package directly, which means that the Go compiler will raise an error if we try to import it normally. But **we need the pq package's init() function to run so that our driver can register itself with database/sql.** We get around this by aliasing the package name to the blank identifier. This means pq.init() still gets executed, but the alias is harmlessly discarded (and our code runs error-free). This approach is standard for most of Go's SQL drivers." - Alex Edwards

## define a book type struct
```
type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}
```
"Next we define a Book type. **The struct fields and their types must align to our books table.** For completeness I should point out that we've only been able to use the string and float32 types safely because we set NOT NULL constraints on the columns in our table. If the table contained nullable fields we would need to use the sql.NullString and sql.NullFloat64 types instead – see [this Gist](https://gist.github.com/alexedwards/dc3145c8e2e6d2fd6cd9) for a working example. Generally it's easiest to avoid nullable fields altogether if you can, which is what we've done here." - Alex Edwards

## initialize a new sql.DB
```
db, err := sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
```
"In the main() function we initialise a new sql.DB by calling sql.Open(). We pass in the name of our driver (in this case "postgres") and the connection string (you'll need to check your driver documentation for the correct format). It's worth emphasizing that **sql.DB is not a database connection – it's an abstraction representing a pool of underlying connections.** You can change the maximum number of open and idle connections in the pool with the db.SetMaxOpenConns() and db.SetMaxIdleConns() methods respectively. A final thing to note is that **sql.DB is safe for concurrent access,** which is very convenient if you're using it in a web application (like we will shortly)." - Alex Edwards

## ping the db
```
	if err = db.Ping(); err != nil {
		panic(err)
	}
```
"Because  sql.Open() doesn't actually check a connection, we also call DB.Ping() to make sure that everything works OK on startup." - Alex Edwards

## query the db
```
rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
```
"We will fetch a resultset from the books table using the DB.Query() method and assign it to a  rows variable. Then we defer rows.Close() to ensure the resultset is properly closed before the parent function returns. **Closing a resultset properly is really important. As long as a resultset is open it will keep the underlying database connection open – which in turn means the connection is not available to the pool.** So if something goes wrong and the resultset isn't closed it can rapidly lead to all the connections in your pool being used up. Another gotcha (which caught me out when I first began) is that the defer statement should come after you check for an error from DB.Query. Otherwise, if DB.Query() returns an error, you'll get a panic trying to close a nil resultset." - Alex Edwards

## iterate through results
```
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}
```
"We then use rows.Next() to iterate through the rows in the resultset. This preps the first (and then each subsequent) row to be acted on by the rows.Scan() method. Note that if iteration over all of the rows completes then the resultset automatically closes itself and frees-up the connection. We use the rows.Scan() method to copy the values from each field in the row to a new Book object that we created. We then check for any errors that occurred during Scan, and add the new Book to a slice of books." - Alex Edwards

## make sure everything ran well
```
	if err = rows.Err(); err != nil {
		panic(err)
	}
```
"When our rows.Next() loop has finished we call rows.Err(). This returns any error that was encountered during the interation. It's important to call this – don't just assume that we completed a successful iteration over the whole resultset." - Alex Edwards Err returns the error, if any, that was encountered during iteration. Err may be called after an explicit or implicit Close.