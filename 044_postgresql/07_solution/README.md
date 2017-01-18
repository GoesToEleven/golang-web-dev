# solution
```
DROP TABLE employees, phonenumbers;
```

```
CREATE TABLE employees (
   ID  SERIAL PRIMARY KEY NOT NULL,
   NAME           TEXT    NOT NULL,
   SCORE            INT     DEFAULT 10 NOT NULL,
   SALARY         REAL
);
```

```
INSERT INTO employees (NAME,SCORE,SALARY) VALUES ('Daniel', 23, 55000.00), ('Arin', 25, 65000.00), ('Juan', 24, 72000.00), ('Shen', 26, 64000.00), ('Myke', 27, 58000.00), ('McLeod', 26, 72000.00), ('James', 32, 35000.00);
```

```
SELECT * FROM employees;
```