# Commands

Here are all of the commands we will use in this section

## log in
```
/Applications/Postgres.app/Contents/Versions/9.6/bin/psql -p5432
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

## log out
```
\q
```

#  create database
```
CREATE DATABASE employees;
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

# insert a record
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

# auto increment key field
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

```
SELECT * FROM phonenumbers;
```

# hands-on exercise
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

# solution
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

# queries
```
SELECT employees.NAME, phonenumbers.PHONE FROM employees INNER JOIN phonenumbers ON employees.ID = phonenumbers.EMP_ID;
```

## cross join

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

# inner join

## two tables

#### review basic select

```
SELECT <fields> FROM <table>;
```

```
SELECT * FROM employees;
```

```
SELECT name, score FROM employees;
```

#### review cross join
```
SELECT <fields> FROM <table1> CROSS JOIN <table2>;
```

```
SELECT person.NAME, sport.NAME FROM person CROSS JOIN sport;
```

### inner join
```
SELECT <fields> FROM <table> INNER JOIN <table>
ON <pkey> = <fkey>;
```

```
SELECT person.NAME, sport.NAME FROM person INNER JOIN sport
ON person.ID = sport.P_ID;
```

```
SELECT employees.NAME, phonenumbers.PHONE FROM employees INNER JOIN phonenumbers ON employees.ID = phonenumbers.EMP_ID;
```

We can use one with our ```people``` and ```sports``` table too, if we wanted, as these tables are connected (remember ```P_ID INT references person(ID)```).

```
SELECT person.NAME, sport.NAME FROM person INNER JOIN sport ON person.ID = sport.P_ID;
```

## three tables - video rental database

#### create a database
```
create database blockbuster;
```

#### switch into the database
```
\c blockbuster
```

#### create three tables

```
create table customers (cid serial primary key not null, cfirst char(50) not null);
```

```
create table movies (mid serial primary key not null, mmovie char(50) not null);
```

```
create table rentals (rid serial primary key not null, cid int references customers(cid), mid int references movies(mid));
```

#### populate tables
```
insert into customers (cfirst) values ('James Bond'), ('Miss Moneypenny'), ('Q'), ('M'), ('Fleming');
```

```
insert into movies (mmovie) values ('Jaws'), ('Alien'), ('Never Say Never'), ('Skyfall'), ('Highlander');
```

```
insert into rentals (cid, mid) values (1,3), (2,5), (4,1), (3,2), (5,4), (3,2), (1,3), (2,4), (5,4), (2,1), (2,3), (4,5), (5,2), (2,1), (3,2), (3,3), (2,3), (1,4), (3,2), (2,3), (3,3), (2,4), (2,3), (1,2), (3,5), (3,4), (1,5);
```

#### inner join query
```
select customers.cfirst, movies.mmovie from customers inner join rentals on customers.cid = rentals.cid inner join movies on rentals.mid = movies.mid;
```


### How this works

```
select * from
tableA inner join tableB
    on tableA.common = tableB.common
inner join TableC
    on tableB.common = TableC.common
```

#### you might also see alias use

```
select * from
tableA a inner join tableB b
        on a.common = b.common
inner join TableC c
        on b.common = c.common
```

# outer join

#### left outer join
  ```
  SELECT person.NAME, sport.NAME FROM person LEFT OUTER JOIN sport ON person.ID = sport.P_ID;
  ```
 
#### right outer join
A right outer join is like a left outer join, but for the table on the right.

```
INSERT INTO sport (NAME) VALUES ('Squirrel Suit Flying');
```

```
  SELECT person.NAME, sport.NAME FROM person RIGHT OUTER JOIN sport ON person.ID = sport.P_ID;
```

#### full outer join
A full outer join is like running both a left outer join and a right outer join at the same time. It gives you everything from all tables, and matches what matches.

```
  SELECT person.NAME, sport.NAME FROM person FULL OUTER JOIN sport ON person.ID = sport.P_ID;
```

# clauses

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

## limit
Limit the number of records returned
```
SELECT * FROM employees LIMIT 4;
```

```
SELECT * FROM employees ORDER BY id LIMIT 4;
```

# update

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

# delete

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

# users & privileges

## see current user
```
SELECT current_user;
```

## details of users
```
\du
```

## create user
```
CREATE USER james WITH PASSWORD 'password';
```

## grant privileges
```
GRANT ALL PRIVILEGES ON DATABASE employees to james;
```

## alter
```
ALTER USER james WITH SUPERUSER;
```

## remove
```
DROP USER james;
```

# Go & postgres

driver
```
go get github.com/lib/pq
```

## create a db

```
CREATE DATABASE bookstore;
```

## create user
```
CREATE USER bond WITH PASSWORD 'password';
```

## grant privileges
```
GRANT ALL PRIVILEGES ON DATABASE bookstore to bond;
```