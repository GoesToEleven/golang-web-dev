CREATE DATABASE exampleunion;

\c exampleunion

CREATE TABLE customers (cid INT PRIMARY KEY NOT NULL, first TEXT);

INSERT INTO customers VALUES (1, 'James'), (2, 'Sergey'), (3, 'Vladimir'), (4, 'Putin'), (5, 'Coup');

CREATE TABLE vendors (vid INT PRIMARY KEY NOT NULL, first TEXT);

INSERT INTO vendors VALUES (1, 'Stacey'), (2, 'Shelley'), (3, 'Sherry'), (4, 'Suzy'), (5, 'Sandy');

CREATE TABLE employees (eid INT PRIMARY KEY NOT NULL, first TEXT);

INSERT INTO employees VALUES (1, 'Jeff'), (2, 'John'), (3, 'Jerry'), (4, 'Jose'), (5, 'Juan');

SELECT v.first FROM vendors AS v UNION SELECT e.first FROM employees AS e;

SELECT 'customer' AS category, c.first FROM customers AS c UNION SELECT 'employee' AS category, e.first FROM employees AS e UNION SELECT 'vendor' AS category, v.first FROM vendors AS v UNION SELECT 'stuffedanimal' AS category, sa.saAnimalName FROM stuffedanimals AS sa;

SELECT 'vendor' AS rowid, v.first FROM vendors AS v UNION SELECT 'employee' AS rowid, e.first FROM employees AS e;

SELECT 'vendor' AS rowid, v.first FROM vendors AS v UNION SELECT 'employee' AS rowid, e.first FROM employees AS e UNION SELECT 'customer' AS rowid, c.first FROM customers AS c;

SELECT 'vendor' AS rowid, v.first FROM vendors AS v UNION SELECT 'employee' AS rowid, e.first FROM employees AS e UNION SELECT 'customer' AS rowid, c.first FROM customers AS c ORDER BY rowid;

SELECT 'vendor' AS rowid, v.first FROM vendors AS v UNION SELECT 'employee' AS rowid, e.first FROM employees AS e UNION SELECT 'customer' AS rowid, c.first FROM customers AS c ORDER BY first;

SELECT 'vendor' AS rowid, v.first FROM vendors AS v UNION SELECT 'employee' AS rowid, e.first FROM employees AS e UNION SELECT 'customer' AS rowid, c.first FROM customers AS c ORDER BY rowid, first;

SELECT 'vendor' AS rowid, v.first AS name FROM vendors AS v UNION SELECT 'employee' AS rowid, e.first AS name FROM employees AS e UNION SELECT 'customer' AS rowid, c.first AS name FROM customers AS c ORDER BY rowid, name;