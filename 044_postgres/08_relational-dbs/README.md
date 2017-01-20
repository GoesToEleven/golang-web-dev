# relational databases

![Understanding relational databases](db.png)

```
CREATE TABLE phonenumbers (
   ID  SERIAL PRIMARY KEY NOT NULL,
   PHONE           CHAR(50) NOT NULL,
   EMP_ID         INT      references employees(ID)
);
```

```
INSERT INTO phonenumbers (PHONE,EMP_ID) VALUES ('555-777-8888', 4), ('555-222-3345', 4), ('777-543-3451', 1), ('544-756-2334', 2);
```