# postgres

[Nice tutorial](https://www.tutorialspoint.com/postgresql/postgresql_create_database.htm)
[documenation reference](https://www.postgresql.org/docs/9.4/static/app-psql.html)

## install
[postgresql website](https://www.postgresql.org/download/)

### mac
Postgres.app is fine.

## log in
```
/Applications/Postgres.app/Contents/Versions/9.6/bin/psql -p5432
```

## log out
```
\q
```

## create database
```
CREATE DATABASE employees;
```

## list databases
```
\l
```

## connect to a database
```
\c <database name>
```

## switch back to postgres database
```
\c postgres
```

## see current user
```
SELECT current_user;
```

## see current database
```
SELECT current_database();
```

## drop (remove, delete) database
```
DROP DATABASE <database name>;
```

## create table
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

Unlike olders, however, schemas can't be nested.

Schemas provide namespacing.

[Read more about schemas](https://www.tutorialspoint.com/postgresql/postgresql_schema.htm)


## insert a record
```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (1, 'Mark', 7, '1212 E. Lane, Someville, AK, 57483', 43000.00 ,'1992-01-13');
```

## list records in a table
```
SELECT * FROM <table name>;
```

## insert a record - variations
omitted values will have the [default value](https://www.postgresql.org/docs/9.3/static/ddl-default.html):
```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,BDAY) VALUES (2, 'Marian', 8, '7214 Wonderlust Ave, Lost Lake, KS, 22897', '1989-11-21');
```

we can use DEFAULT rather leaving a field blank or specifying a value:
```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (3, 'Maxwell', 6, '7215 Jasmine Place, Corinda, CA 98743', 87500.00, DEFAULT);
```

we can insert multiple rows:
```
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (4, 'Jasmine', 5, '983 Star Ave., Brooklyn, NY, 00912 ', 55700.00, '1997-12-13' ), (5, 'Orranda', 9, '745 Hammer Lane, Hammerfield, Texas, 75839', 65350.00 , '1992-12-13');
```

## auto increment key field
Instead of creating a unique ID number ourselves, we can have postgres automatically increment this ID field.
 
 To do this we use the data types smallserial, serial or bigserial (not true types but for convenience).
 
 This is like AUTO_INCREMENT in other databases.

```
CREATE TABLE phonenumbers(
	ID  SERIAL PRIMARY KEY,
	PHONE           TEXT      NOT NULL
);
```

```
INSERT INTO phonenumbers (PHONE) VALUES ( '234-432-5234'), ('543-534-6543'), ('312-123-5432');
```

```
\d phonenumbers
```

## hands-on exercise
1. delete all of your current tables.
1. READ ALL OF THIS: create a new table called employees with these fields ```id, name, score, salary``` AND give ```score``` a default value of 10 AND have the ```id``` field automatically increment.
1. add these records and then show all of the records
```
 id |  name  | score | salary 
----+--------+-------+--------
  1 | Daniel |    23 |  55000
  2 | Arin   |    25 |  65000
  3 | Juan   |    24 |  72000
  4 | Shen   |    26 |  64000
  5 | Myke   |    27 |  58000
  6 | McLeod |    26 |  72000
  7 | James  |    32 |  35000
```

## solution
```
DROP TABLE employees, phonenumbers;
```

```
CREATE TABLE employees (
   ID  SERIAL PRIMARY KEY NOT NULL,
   NAME           TEXT    NOT NULL,
   SCORE            INT     DEFAULT 10 NOT NULL,
   SALARY         REAL
);
```

```
INSERT INTO employees (NAME,SCORE,SALARY) VALUES ('Daniel', 23, 55000.00), ('Arin', 25, 65000.00), ('Juan', 24, 72000.00), ('Shen', 26, 64000.00), ('Myke', 27, 58000.00), ('McLeod', 26, 72000.00), ('James', 32, 35000.00);
```

```
SELECT * FROM employees;
```

## relational databases

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

/////////////

# Organizing code

## create table
We will be following the example created in [this article by Alex Edwards](http://www.alexedwards.net/blog/practical-persistence-sql) and licensed under a [MIT license](https://opensource.org/licenses/MIT)

```
CREATE TABLE books (
  isbn    char(14) NOT NULL,
  title   varchar(255) NOT NULL,
  author  varchar(255) NOT NULL,
  price   decimal(5,2) NOT NULL
);

## insert records
```
INSERT INTO books (isbn, title, author, price) VALUES
('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);
```

##
```
ALTER TABLE books ADD PRIMARY KEY (isbn);
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

##
```
```