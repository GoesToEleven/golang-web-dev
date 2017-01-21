# auto increment key field
Instead of creating a unique ID number ourselves, we can have postgres automatically increment this ID field.
 
 To do this we use the data types smallserial, serial or bigserial (not true types but for convenience).
 
 This is like AUTO_INCREMENT in other databases.

```
CREATE TABLE phonenumbers(
	ID  SERIAL PRIMARY KEY,
	PHONE           TEXT      NOT NULL
);
```

```
INSERT INTO phonenumbers (PHONE) VALUES ( '234-432-5234'), ('543-534-6543'), ('312-123-5432');
```

```
\d phonenumbers
```

```
SELECT * FROM phonenumbers;
```