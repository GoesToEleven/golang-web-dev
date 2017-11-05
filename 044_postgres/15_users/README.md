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
privileges: SELECT, INSERT, UPDATE, DELETE, RULE, ALL
```
GRANT ALL PRIVILEGES ON DATABASE company to james;
```

## revoke privileges
```
REVOKE ALL PRIVILEGES ON DATABASE company from james;
```

## alter
```
ALTER USER james WITH SUPERUSER;
```

```
ALTER USER james WITH NOSUPERUSER;
```

## remove
```
DROP USER james;
```
