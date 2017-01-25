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