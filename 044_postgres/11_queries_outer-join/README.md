# outer join

## left outer join
 A left outer join gives you everything in one table, and also the matching records in another table.
 
 For tables A and B a left outer join would give you all rows of the "left" table (A), even if the join-condition does not find any matching row in the "right" table (B).
  
 This means that if the ON clause matches 0 (zero) rows in B (for a given row in A), the join will still return a row in the result (for that row)—but with NULL in each column from B.
 
  ```
  SELECT person.NAME, sport.NAME FROM person LEFT OUTER JOIN sport ON person.ID = sport.P_ID;
  ```
 
## right outer join
A right outer join is like a left outer join, but for the table on the right.

```
INSERT INTO sport (NAME) VALUES ('Squirrel Suit Flying');
```

```
  SELECT person.NAME, sport.NAME FROM person RIGHT OUTER JOIN sport ON person.ID = sport.P_ID;
```

## full outer join
A full outer join is like running both a left outer join and a right outer join at the same time. It gives you everything from all tables, and matches what matches.

```
  SELECT person.NAME, sport.NAME FROM person FULL OUTER JOIN sport ON person.ID = sport.P_ID;
```