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
CREATE USER bond WITH PASSWORD 'password';
```

## grant privileges
```
GRANT ALL PRIVILEGES ON DATABASE employees to bond;
```

## alter
```
ALTER USER bond WITH SUPERUSER;
```

## remove
```
DROP USER bond;
```