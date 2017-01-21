# create three tables

```
CREATE TABLE people (
   ID  SERIAL PRIMARY KEY NOT NULL,
   NAME           CHAR(50) NOT NULL
);
```

```
CREATE TABLE videos (
   ID  SERIAL PRIMARY KEY NOT NULL,
   NAME           CHAR(50) NOT NULL
);
```

```
CREATE TABLE rentals (
   ID  SERIAL PRIMARY KEY NOT NULL,
   V_ID         INT      references people(ID),
   P_ID         INT      references videos(ID)
);
```

# populate tables
```
INSERT INTO people (NAME) VALUES ('Shen'), ('Daniel'), ('Juan'), ('Arin'), ('McLeod');
```

```
INSERT INTO videos (NAME) VALUES ('Jaws'), ('Alien'), ('Never Say Never'), ('Skyfall'), ('Highlander');
```

```
INSERT INTO rentals (V_ID, P_ID) VALUES (1,3), (2,5), (4,1), (3,2), (5,4), (3,2), (1,3), (2,4), (5,4), (2,1), (2,3), (4,5), (5,2), (2,1), (3,2), (3,3), (2,3), (1,4), (3,2), (2,3), (3,3), (2,4), (2,3), (1,2), (3,5), (3,4), (1,5);
```

# inner join query
```
SELECT people.NAME, videos.NAME FROM
people INNER JOIN rentals
    ON people.ID = rentals.P_ID
INNER JOIN videos
    ON rentals.V_ID = videos.ID;
```


# How this works

```
select * from
tableA inner join tableB
    on tableA.common = tableB.common
inner join TableC
    on tableB.common = TableC.common
```

## you might also see alias use

```
select * from
tableA a inner join tableB b
        on a.common = b.common
inner join TableC c
        on b.common = c.common
```











