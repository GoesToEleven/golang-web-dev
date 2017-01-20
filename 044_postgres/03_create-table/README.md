# create table
```
CREATE TABLE employees (
   ID INT PRIMARY KEY     NOT NULL,
   NAME           TEXT    NOT NULL,
   RANK           INT     NOT NULL,
   ADDRESS        CHAR(50),
   SALARY         REAL DEFAULT 25500.00,
   BDAY			  DATE DEFAULT '1900-01-01'
);
```

## show tables in a database (list down)
```
\d
```

## show details of a table
```
\d <table name>
```

## drop a table
```
DROP TABLE <table name>;
```

## schema
Schemas allow us to organize our database and database code.

A schema is like a folder.

Into this folder, you can put tables, views, indexes, sequences, data types, operators, and functions. 

Unlike folders, however, schemas can't be nested.

Schemas provide namespacing.

[Read more about schemas](https://www.tutorialspoint.com/postgresql/postgresql_schema.htm)