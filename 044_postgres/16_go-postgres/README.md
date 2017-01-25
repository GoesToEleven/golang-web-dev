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
