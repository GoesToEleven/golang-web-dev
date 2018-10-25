CREATE DATABASE sales;

\l

\c sales

CREATE TABLE employees (eid INT PRIMARY KEY NOT NULL, ename TEXT NOT NULL);

CREATE TABLE states (stid INT PRIMARY KEY NOT NULL, stname TEXT NOT NULL);

CREATE TABLE cities (cid INT PRIMARY KEY NOT NULL, cname TEXT NOT NULL);

CREATE TABLE sales (sid INT PRIMARY KEY NOT NULL, eid INT REFERENCES employees(eid) NOT NULL, cid INT REFERENCES cities(cid) NOT NULL, stid INT REFERENCES states(stid) NOT NULL, samount INT NOT NULL);

INSERT INTO employees VALUES (1,'James'), (2,'Moneypenny'), (3,'Ian'), (4,'Fleming'), (5,'Bond');

INSERT INTO states VALUES (1,'AL'), (2,'AK'), (3,'AZ'), (4,'AR'), (5,'CA');

INSERT INTO cities VALUES (1,'Montgomery'), (2,'Anchorage'), (3,'Phoenix'), (4,'Little Rock'), (5,'Sacramento'), (6,'Fresno'), (7,'Benicia'), (8,'San Fran'), (9,'LA'), (10,'San Diego');

https://play.golang.org/p/knG5I2AAi79

INSERT INTO sales VALUES (1,3,2,2,1948), (2,2,10,5,1419), (3,1,6,5,557), (4,5,1,1,3612), (5,5,3,3,4829), (6,2,5,5,1546), (7,2,8,5,596), (8,4,7,5,1359), (9,3,8,5,3388), (10,1,9,5,3116), (11,4,2,2,2488), (12,5,2,2,457), (13,2,8,5,1586), (14,4,7,5,3191), (15,4,5,5,2534), (16,4,8,5,4425), (17,4,10,5,2058), (18,5,2,2,2300), (19,1,1,1,2989), (20,4,9,5,4456), (21,1,2,2,2706), (22,2,7,5,4929), (23,3,2,2,4884), (24,4,7,5,4477), (25,4,3,3,548), (26,3,5,5,2564), (27,1,7,5,3724), (28,3,4,4,3234), (29,5,2,2,3134), (30,2,4,4,2103), (31,2,9,5,2647), (32,1,8,5,1604), (33,4,3,3,2306), (34,1,9,5,1452), (35,3,6,5,3788), (36,1,1,1,386), (37,3,1,1,3199), (38,2,4,4,3683), (39,3,5,5,4368), (40,2,8,5,995), (41,3,7,5,4082), (42,2,10,5,1371), (43,2,4,4,4920), (44,3,2,2,2276), (45,1,6,5,1488), (46,4,10,5,2919), (47,4,5,5,1325), (48,3,8,5,1633), (49,5,7,5,641), (50,2,7,5,3177), (51,2,1,1,3945), (52,1,5,5,4284), (53,1,2,2,1703), (54,3,9,5,3332), (55,5,9,5,2923), (56,3,4,4,4309), (57,4,4,4,1267), (58,1,1,1,541), (59,3,5,5,4758), (60,2,6,5,3140), (61,4,1,1,4801), (62,1,10,5,884), (63,5,1,1,3348), (64,1,1,1,4263), (65,1,10,5,2149), (66,1,7,5,3767), (67,2,1,1,1730), (68,2,3,3,2678), (69,1,7,5,500), (70,3,3,3,393), (71,2,9,5,4204), (72,4,9,5,3857), (73,3,10,5,1258), (74,1,3,3,2282), (75,1,4,4,518), (76,1,4,4,4097), (77,3,7,5,2621), (78,3,1,1,2130), (79,1,2,2,2521), (80,5,10,5,1565), (81,2,1,1,1282), (82,2,8,5,701), (83,3,10,5,1634), (84,5,2,2,1786), (85,1,10,5,1820), (86,1,5,5,3763), (87,1,1,1,4385), (88,1,4,4,1544), (89,1,2,2,675), (90,1,10,5,1437), (91,3,9,5,1645), (92,5,6,5,1338), (93,4,5,5,1707), (94,3,9,5,2521), (95,2,3,3,1218), (96,1,2,2,3783), (97,3,10,5,230), (98,2,2,2,956), (99,4,2,2,1154), (100,5,7,5,4968);

SELECT e.ename, c.cname, st.stname, s.samount FROM ((employees AS e JOIN sales AS s ON e.eid = s.eid) JOIN cities AS c ON s.cid = c.cid) JOIN states AS st ON s.stid = st.stid;



SELECT e.ename, c.cname, st.stname, s.samount FROM ((employees AS e JOIN sales AS s ON e.eid = s.eid) JOIN cities AS c ON s.cid = c.cid) JOIN states AS st ON s.stid = st.stid WHERE e.eid = 1;


// does not work
SELECT e.ename, c.cname, st.stname, s.samount FROM (((employees AS e WHERE e.eid = 1) JOIN sales AS s ON e.eid = s.eid) JOIN cities AS c ON s.cid = c.cid) JOIN states AS st ON s.stid = st.stid;

// works
SELECT e.ename FROM employees AS e WHERE e.eid = 1;

SELECT e.ename, c.cname, st.stname, SUM(s.samount) FROM ((employees AS e JOIN sales AS s ON e.eid = s.eid) JOIN cities AS c ON s.cid = c.cid) JOIN states AS st ON s.stid = st.stid GROUP BY e.ename, c.cname, st.stname;

SELECT e.ename, c.cname, st.stname, SUM(s.samount) FROM ((employees AS e JOIN sales AS s ON e.eid = s.eid) JOIN cities AS c ON s.cid = c.cid) JOIN states AS st ON s.stid = st.stid WHERE e.eid = 1 GROUP BY e.ename, c.cname, st.stname;

SELECT e.ename, c.cname, st.stname, SUM(s.samount) FROM ((employees AS e JOIN sales AS s ON e.eid = s.eid) JOIN cities AS c ON s.cid = c.cid) JOIN states AS st ON s.stid = st.stid GROUP BY e.ename, c.cname, st.stname ORDER BY e.ename;

SELECT e.ename, c.cname, st.stname, SUM(s.samount) FROM ((employees AS e JOIN sales AS s ON e.eid = s.eid) JOIN cities AS c ON s.cid = c.cid) JOIN states AS st ON s.stid = st.stid WHERE e.eid = 1 GROUP BY e.ename, c.cname, st.stname HAVING st.stname = 'CA';

SELECT e.ename, c.cname, st.stname, SUM(s.samount) FROM ((employees AS e JOIN sales AS s ON e.eid = s.eid) JOIN cities AS c ON s.cid = c.cid) JOIN states AS st ON s.stid = st.stid WHERE e.eid = 1 AND st.stname = 'CA' GROUP BY e.ename, c.cname, st.stname;


SELECT e.ename, c.cname, st.stname, SUM(s.samount) FROM ((employees AS e JOIN sales AS s ON e.eid = s.eid) JOIN cities AS c ON s.cid = c.cid) JOIN states AS st ON s.stid = st.stid GROUP BY e.ename, c.cname, st.stname HAVING e.ename = 'James' AND st.stname = 'CA';

// does not work
SELECT e.ename, c.cname, st.stname, SUM(s.samount) FROM ((employees AS e JOIN sales AS s ON e.eid = s.eid) JOIN cities AS c ON s.cid = c.cid) JOIN states AS st ON s.stid = st.stid GROUP BY e.ename, c.cname, st.stname HAVING e.eid = 1 AND st.stname = 'CA';


