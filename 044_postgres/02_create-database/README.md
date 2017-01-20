#  create database
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