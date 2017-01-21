# create a database
```
create database blockbuster;
```

# switch into the database
```
\c blockbuster
```

# create three tables

```
create table customers (cid serial primary key not null, cfirst char(50) not null);
```

```
create table movies (mid serial primary key not null, mmovie char(50) not null);
```

```
create table rentals (rid serial primary key not null, cid int references customers(cid), mid int references movies(mid));
```

# populate tables
```
insert into customers (cfirst) values ('James Bond'), ('Miss Moneypenny'), ('Q'), ('M'), ('Fleming');
```

```
insert into movies (mmovie) values ('Jaws'), ('Alien'), ('Never Say Never'), ('Skyfall'), ('Highlander');
```

```
insert into rentals (cid, mid) values (1,3), (2,5), (4,1), (3,2), (5,4), (3,2), (1,3), (2,4), (5,4), (2,1), (2,3), (4,5), (5,2), (2,1), (3,2), (3,3), (2,3), (1,4), (3,2), (2,3), (3,3), (2,4), (2,3), (1,2), (3,5), (3,4), (1,5);
```

# inner join query
```
select customers.cfirst, movies.mmovie from customers inner join rentals on customers.cid = rentals.cid inner join movies on rentals.mid = movies.mid;
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











