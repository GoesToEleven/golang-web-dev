## where
Adding **WHERE** to a SQL query allows you to filter results.
```
SELECT * FROM employees WHERE salary > 60000;
```

## and
```
SELECT * FROM employees WHERE salary > 60000 AND score = 26;
```

## in
```
SELECT * FROM employees WHERE score IN (25,26);
```

## not
```
SELECT * FROM employees WHERE score NOT IN (25,26);
```

## between
```
SELECT * FROM employees WHERE score BETWEEN 23 AND 26;
```

## is not null
```
SELECT * FROM employees WHERE score IS NOT NULL;
```

## like
```
SELECT * FROM employees WHERE name LIKE '%an%';
```

## or
```
SELECT * FROM employees WHERE score <= 24 OR salary < 50000;
```

## limit
Limit the number of records returned
```
SELECT * FROM employees LIMIT 4;
```

```
SELECT * FROM employees ORDER BY id LIMIT 4;
```