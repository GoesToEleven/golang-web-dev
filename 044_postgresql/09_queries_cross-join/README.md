# queries

## querying from more than one table

This query is an **inner join**.
```
SELECT employees.NAME, phonenumbers.PHONE FROM employees INNER JOIN phonenumbers ON employees.ID = phonenumbers.EMP_ID;
```

# join queries

Join queries allow us to select records from two or more tables.

A join query combines columns from one or more tables - it joins a bunch of columns from different tables together.

There are five types of joins in postgres:

## cross join
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