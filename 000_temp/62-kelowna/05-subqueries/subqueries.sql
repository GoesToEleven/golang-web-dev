CREATE DATABASE acmesales;

\c acmesales

CREATE TABLE customers (cid INT PRIMARY KEY NOT NULL, cfirst TEXT);

INSERT INTO customers VALUES (1,'Cory'), (2,'Casey'), (3,'Canard'), (4,'Cully'), (5,'Coffer'), (6,'Homey');

CREATE TABLE items (iid INT PRIMARY KEY NOT NULL, iitem TEXT);

INSERT INTO items VALUES (1,'dog food'), (2,'tiramisu'), (3,'vino'), (4,'cheese'), (5,'chocolate'), (6,'gummy worms'), (7,'mango');

CREATE TABLE orders (oid INT PRIMARY KEY NOT NULL, cid INT REFERENCES customers(cid), odate TEXT);

INSERT INTO orders VALUES (1,2,'tuesday'), (2,5,'wednesday'), (3,1,'thursday'), (4,3,'sunday'), (5,4,'saturday'), (6,5,'friday');

CREATE TABLE orderitems (oiid INT PRIMARY KEY NOT NULL, oid INT REFERENCES orders(oid), iid INT REFERENCES items(iid));

INSERT INTO orderitems VALUES (1,4,6), (2,4,5), (3,5,4), (4,6,3), (5,1,2), (6,2,1), (7,3,6), (8,4,5), (9,5,4), (10,6,3);

SELECT c.cfirst, o.oid, o.odate, i.iitem FROM customers AS c JOIN orders AS o ON c.cid = o.cid JOIN orderitems AS oi ON o.oid = oi.oid JOIN items AS i ON oi.iid = i.iid;

SELECT c.cfirst, o.oid, o.odate, i.iitem FROM customers AS c FULL JOIN orders AS o ON c.cid = o.cid FULL JOIN orderitems AS oi ON o.oid = oi.oid FULL JOIN items AS i ON oi.iid = i.iid;

// SCALAR
// an expression that evaluates to one value
// aka, one column & one row = one field value

SELECT o.oid, o.odate, (SELECT c.cfirst FROM customers AS c WHERE c.cid = o.cid) FROM orders AS o;

// could have also done the above this way

SELECT o.oid, o.odate, c.cfirst FROM customers AS c JOIN orders AS o ON c.cid = o.cid;

// aggregate function COUNT

SELECT COUNT(cfirst) AS ourcustomers FROM customers;

SELECT c.cfirst, (SELECT COUNT(*) FROM orders AS o WHERE c.cid = o.cid) AS ordercount FROM customers AS c;

// aggregate function MAX

SELECT MAX(cid) AS highestcid FROM customers;

SELECT MAX(odate) AS maxday FROM orders;

SELECT c.cfirst, (SELECT MAX(o.oid) FROM orders AS o WHERE c.cid = o.cid) AS maxoid FROM customers AS c;

// subquery as filters
// subquery with the WHERE clause
SELECT o.oid, o.odate, c.cfirst FROM customers AS c JOIN orders AS o ON c.cid = o.cid WHERE o.odate = (SELECT MAX(odate) FROM orders);

// subquery as filters
// subquery with the IN keyword
SELECT o.oid, o.odate, c.cfirst FROM customers AS c JOIN orders AS o ON c.cid = o.cid WHERE o.odate IN (SELECT odate FROM orders WHERE odate LIKE '%u%');

// ALL, SOME, ANY keywords
SELECT o.oid, o.odate, c.cfirst FROM customers AS c JOIN orders AS o ON c.cid = o.cid WHERE o.odate = ANY (SELECT odate FROM orders WHERE odate LIKE '%u%');

SELECT o.oid, o.odate, c.cfirst FROM customers AS c JOIN orders AS o ON c.cid = o.cid WHERE o.odate = SOME (SELECT odate FROM orders WHERE odate LIKE '%u%');

SELECT o.oid, o.odate, c.cfirst FROM customers AS c JOIN orders AS o ON c.cid = o.cid WHERE o.odate = ALL (SELECT odate FROM orders WHERE odate LIKE '%u%');

SELECT o.oid, o.odate, c.cfirst FROM customers AS c JOIN orders AS o ON c.cid = o.cid WHERE o.odate = ALL (SELECT odate FROM orders WHERE odate LIKE 'thur%');

SELECT o.oid, o.odate, c.cfirst FROM customers AS c JOIN orders AS o ON c.cid = o.cid WHERE o.odate > ALL (SELECT odate FROM orders WHERE odate LIKE 'thur%');

// EXISTS
// I DO NOT KNOW HOW THIS WORKS
// MORE RESEARCH AND EXPERIMENTATION ARE NEEDED
SELECT c.cfirst, o.oid, o.odate, i.iitem FROM customers AS c JOIN orders AS o ON c.cid = o.cid JOIN orderitems AS oi ON o.oid = oi.oid JOIN items AS i ON oi.iid = i.iid;

SELECT c.cfirst, o.oid, o.odate, i.iitem FROM customers AS c JOIN orders AS o ON c.cid = o.cid JOIN orderitems AS oi ON o.oid = oi.oid JOIN items AS i ON oi.iid = i.iid WHERE i.iitem = 'chocolate';

SELECT c.cfirst, o.oid, o.odate, i.iitem FROM customers AS c JOIN orders AS o ON c.cid = o.cid JOIN orderitems AS oi ON o.oid = oi.oid JOIN items AS i ON oi.iid = i.iid WHERE EXISTS (SELECT c.cfirst, o.oid, o.odate, i.iitem FROM customers AS c JOIN orders AS o ON c.cid = o.cid JOIN orderitems AS oi ON o.oid = oi.oid JOIN items AS i ON oi.iid = i.iid WHERE i.iitem = 'chocolate');




