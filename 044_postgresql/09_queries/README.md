# relational databases

![Understanding relational databases](db.png)

```
CREATE TABLE phonenumbers (
   ID  SERIAL PRIMARY KEY NOT NULL,
   PHONE           CHAR(50) NOT NULL,
   EMP_ID         INT      references employees(ID)
);
```

```
INSERT INTO phonenumbers (PHONE,EMP_ID) VALUES ('555-777-8888', 4), ('555-222-3345', 4), ('777-543-3451', 1), ('544-756-2334', 2);
```

## querying from more than one table

```
SELECT employees.NAME, phonenumbers.PHONE FROM employees INNER JOIN phonenumbers ON employees.ID = phonenumbers.EMP_ID;
```

## join queries

Join queries allow us to select records from two or more tables.

A join query combines columns from one or more tables - it joins a bunch of columns from different tables together.

There are five types of joins in postgres:

1. CROSS JOIN
A cross join returns the **Cartesian product** of rows from tables in the join. In other words, it will produce rows which combine each row from the first table with each row from the second table.

```
CREATE TABLE person (
   ID  SERIAL PRIMARY KEY NOT NULL,
   NAME           CHAR(50) NOT NULL
);
```

```
INSERT INTO person (NAME) VALUES ('Shen'), ('Daniel'), ('Juan'), ('Arin'), ('McLeod');
```

```
CREATE TABLE sport (
   ID  SERIAL PRIMARY KEY NOT NULL,
   NAME           CHAR(50) NOT NULL,
   P_ID         INT      references person(ID)
);
```

```
INSERT INTO sport (NAME, P_ID) VALUES ('Surf',1),('Soccer',3),('Ski',3),('Sail',3),('Bike',3);
```

```
SELECT person.NAME, sport.NAME FROM person CROSS JOIN sport;
```

1. INNER JOIN
An inner join allows us to select records from two tables. 
 
We used an inner join above when we asked for the phone numbers associated with an employee:

```
SELECT employees.NAME, phonenumbers.PHONE FROM employees INNER JOIN phonenumbers ON employees.ID = phonenumbers.EMP_ID;
```

We can use one with our ```people``` and ```sports``` table too, if we wanted, as these tables are connected (remember ```P_ID INT references person(ID)```).

```
SELECT person.NAME, sport.NAME FROM person INNER JOIN sport ON person.ID = sport.P_ID;
```

Here is how wikipedia explains an inner join:

An inner join requires each row in the two joined tables to have matching rows, and is a commonly used join operation in applications but should not be assumed to be the best choice in all situations. 

Inner join creates a new result table by combining column values of two tables (A and B) based upon the join-predicate. 

The query compares each row of A with each row of B to find all pairs of rows which satisfy the join-predicate. 

When the join-predicate is satisfied by matching non-NULL values, column values for each matched pair of rows of A and B are combined into a result row. 

The result of the join can be defined as the outcome of first taking the Cartesian product (or Cross join) of all rows in the tables (combining every row in table A with every row in table B) and then returning all rows which satisfy the join predicate. Actual SQL implementations normally use other approaches, such as hash joins or sort-merge joins, since computing the Cartesian product is slower and would often require a prohibitively large amount of memory to store. 

SQL specifies two different syntactical ways to express joins: the "explicit join notation" and the "implicit join notation". The "implicit join notation" is no longer considered a best practice, although database systems still support it. The **"explicit join notation"** uses the JOIN keyword, optionally preceded by the INNER keyword, to specify the table to join, and the ON keyword to specify the predicates for the join.

1. LEFT OUTER JOIN
 A left outer join gives you everything in one table, and also the matching records in another table.
 
 For tables A and B a left outer join would give you all rows of the "left" table (A), even if the join-condition does not find any matching row in the "right" table (B).
  
 This means that if the ON clause matches 0 (zero) rows in B (for a given row in A), the join will still return a row in the result (for that row)—but with NULL in each column from B.
 
  ```
  SELECT person.NAME, sport.NAME FROM person LEFT OUTER JOIN sport ON person.ID = sport.P_ID;
  ```
 
1. RIGHT OUTER JOIN
A right outer join is like a left outer join, but for the table on the right.

```
INSERT INTO sport (NAME) VALUES ('Squirrel Suit Flying');
```

```
  SELECT person.NAME, sport.NAME FROM person RIGHT OUTER JOIN sport ON person.ID = sport.P_ID;
```

1. FULL OUTER JOIN
A full outer join is like running both a left outer join and a right outer join at the same time. It gives you everything from all tables, and matches what matches.

```
  SELECT person.NAME, sport.NAME FROM person FULL OUTER JOIN sport ON person.ID = sport.P_ID;
```

## where

Adding **WHERE** to a SQL query allows you to filter results.

```
SELECT * FROM employees WHERE salary > 60000;
```

### and
```
SELECT * FROM employees WHERE salary > 60000 AND score = 26;
```

### in
```
SELECT * FROM employees WHERE score IN (25,26);
```

### not
```
SELECT * FROM employees WHERE score NOT IN (25,26);
```

### between
```
SELECT * FROM employees WHERE score BETWEEN 23 AND 26;
```

### is not null
```
SELECT * FROM employees WHERE score IS NOT NULL;
```

### like
```
SELECT * FROM employees WHERE name LIKE '%an%';
```

### or
```
SELECT * FROM employees WHERE score <= 24 OR salary < 50000;
```

## update

syntax
```
UPDATE table
SET col1 = val1, col2 = val2, ..., colN = valN
WHERE <condition>;
```

```
SELECT * FROM employees;
```

```
UPDATE employees SET score = 99 WHERE ID = 3;
```

## order by
```
SELECT * FROM employees ORDER BY id;
```

## delete

syntax
```
DELETE FROM table
WHERE <condition>;
```

```
SELECT * FROM sport;
```

```
DELETE FROM sport WHERE id = 6;
```

**WARNING: this deletes all records:**
```
DELETE FROM sport;
```

## limit

Limit the number of records returned

```
SELECT * FROM employees LIMIT 4;
```

```
SELECT * FROM employees ORDER BY id LIMIT 4;
```
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