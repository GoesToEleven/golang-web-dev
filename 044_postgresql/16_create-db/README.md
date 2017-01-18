# Go & postgres

driver
```
go get github.com/lib/pq
```

## create database

We will be following the example created in [this article by Alex Edwards](http://www.alexedwards.net/blog/practical-persistence-sql) and licensed under a [MIT license](https://opensource.org/licenses/MIT)

```
CREATE DATABASE bookstore;
```

list databases
```
\l
```

## create table
```
CREATE TABLE books (
  isbn    char(14)     PRIMARY KEY NOT NULL,
  title   varchar(255) NOT NULL,
  author  varchar(255) NOT NULL,
  price   decimal(5,2) NOT NULL
);
```

list tabls
```
\d
```

list table details
```
\d books
```

## insert records
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


/////////////

# Organizing code




##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```

##
```
```