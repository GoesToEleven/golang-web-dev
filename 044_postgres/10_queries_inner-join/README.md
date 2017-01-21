# general syntax

## select
```
SELECT <fields> FROM <table>;
```

```
SELECT * FROM employees;
```

```
SELECT name, score FROM employees;
```

## cross join
```
SELECT <fields> FROM <table1> CROSS JOIN <table2>;
```

```
SELECT person.NAME, sport.NAME FROM person CROSS JOIN sport;
```

## inner join
```
SELECT <fields> FROM <table> INNER JOIN <table>
ON <pkey> = <fkey>;
```

```
SELECT person.NAME, sport.NAME FROM person INNER JOIN sport
ON person.ID = sport.P_ID;
```

# inner join

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