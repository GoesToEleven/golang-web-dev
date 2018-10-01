/* COUNT */

/* from cooking database */

SELECT COUNT(recipetitle) AS totalRecipes FROM recipes;

SELECT COUNT(DISTINCT recipetitle) AS totalDistRecipes FROM recipes;

/* from hollywood database*/
SELECT c.cname, m.mname, a.aname FROM ((((customers AS c JOIN rentals AS r ON c.cid = r.cid) JOIN movies AS m ON r.mid = m.mid) JOIN castmembers AS cm ON m.mid = cm.mid) JOIN actors AS a ON cm.aid = a.aid);

SELECT c.cname, m.mname, a.aname FROM ((((customers AS c LEFT JOIN rentals AS r ON c.cid = r.cid) LEFT JOIN movies AS m ON r.mid = m.mid) LEFT JOIN castmembers AS cm ON m.mid = cm.mid) LEFT JOIN actors AS a ON cm.aid = a.aid);

SELECT r.rid, m.mname FROM rentals AS r JOIN movies AS m ON r.mid = m.mid;

SELECT r.rid, m.mname FROM rentals AS r RIGHT JOIN movies AS m ON r.mid = m.mid;

SELECT cm.caid, a.aname FROM castmembers AS cm RIGHT JOIN actors AS a ON cm.aid = a.aid;

SELECT cm.caid, a.aname FROM actors AS a LEFT JOIN castmembers AS cm ON a.aid = cm.aid;

SELECT c.cname, m.mname, a.aname FROM ((((customers AS c FULL JOIN rentals AS r ON c.cid = r.cid) FULL JOIN movies AS m ON r.mid = m.mid) FULL JOIN castmembers AS cm ON m.mid = cm.mid) FULL JOIN actors AS a ON cm.aid = a.aid);

SELECT COUNT(*) AS totalRentals FROM rentals;

SELECT COUNT(cname) AS custs FROM customers;

SELECT COUNT(DISTINCT cname) AS distinctCusts FROM customers;



/* SUM */

DROP TABLE rentals;

CREATE TABLE rentals (rid INT PRIMARY KEY NOT NULL, cid INT NOT NULL REFERENCES customers(cid), mid INT NOT NULL REFERENCES movies(mid), rAmount INT NOT NULL);

https://play.golang.org/p/H7afRPh5bUd

INSERT INTO rentals VALUES (1,2,8,5),(2,4,2,6),(3,8,1,4),(4,7,5,4),(5,11,10,6),(6,7,2,3),(7,9,7,3),(8,1,9,6),(9,10,8,5),(10,7,1,3),(11,1,9,5),(12,2,10,4),(13,2,2,3),(14,5,4,3),(15,6,4,6),(16,5,9,7),(17,6,4,5),(18,8,10,7),(19,8,6,6),(20,11,4,3),(21,5,1,3),(22,3,7,6),(23,2,3,6),(24,4,4,4),(25,6,9,5),(26,1,8,6),(27,7,1,6),(28,11,8,6),(29,11,10,6),(30,9,2,5),(31,4,7,4),(32,11,1,6),(33,4,4,3),(34,10,6,4),(35,1,8,5),(36,4,1,3),(37,7,3,6),(38,4,2,5),(39,11,8,5),(40,8,2,7),(41,4,3,4),(42,10,7,3),(43,1,7,7),(44,6,3,3),(45,10,1,5),(46,7,9,6),(47,1,4,7),(48,4,3,5),(49,7,10,3),(50,5,2,4),(51,7,2,7),(52,10,6,6),(53,4,1,5),(54,3,8,4),(55,7,5,5),(56,7,3,6),(57,8,9,4),(58,5,6,3),(59,8,3,5),(60,5,2,7),(61,3,4,3),(62,7,1,6),(63,1,5,5),(64,3,6,5),(65,10,1,6),(66,7,6,4),(67,11,7,7),(68,8,2,5),(69,5,1,7),(70,5,8,5),(71,8,2,6),(72,1,9,4),(73,5,8,5),(74,8,6,4),(75,11,6,5),(76,9,1,4),(77,11,3,3),(78,10,3,7),(79,5,1,3),(80,7,5,7),(81,10,2,4),(82,2,7,3),(83,3,8,6),(84,6,5,3),(85,1,6,7),(86,5,1,5),(87,9,1,7),(88,6,6,6),(89,7,1,7),(90,3,1,4),(91,11,3,7),(92,2,5,5),(93,4,4,4),(94,5,3,3),(95,8,2,5),(96,9,6,5),(97,10,3,7),(98,10,2,3),(99,9,4,6),(100,8,10,5);

SELECT c.cname, m.mname, a.aname, r.rAmount FROM ((((customers AS c FULL JOIN rentals AS r ON c.cid = r.cid) FULL JOIN movies AS m ON r.mid = m.mid) FULL JOIN castmembers AS cm ON m.mid = cm.mid) FULL JOIN actors AS a ON cm.aid = a.aid);

SELECT SUM(rAmount) AS totalRev FROM rentals;

select c.cname, r.ramount from customers AS c JOIN rentals AS r ON c.cid = r.cid;

select c.cname, r.rid, r.ramount from customers AS c JOIN rentals AS r ON c.cid = r.cid;

select c.cname, r.rid, r.ramount from customers AS c FULL JOIN rentals AS r ON c.cid = r.cid;

SELECT c.cid, c.cname, r.rid, r.ramount FROM customers AS c FULL JOIN rentals AS r ON c.cid = r.cid;

SELECT c.cid, c.cname, r.rid, r.ramount, m.mname, a.aname FROM ((((customers AS c FULL JOIN rentals AS r ON c.cid = r.cid) FULL JOIN movies AS m ON r.mid = m.mid) FULL JOIN castmembers AS cm ON m.mid = cm.mid) FULL JOIN actors AS a ON cm.aid = a.aid);

SELECT c.cid, c.cname, SUM(r.ramount) AS bigSpend FROM customers AS c FULL JOIN rentals AS r ON c.cid = r.cid GROUP BY c.cid;

SELECT c.cid, c.cname, SUM(r.ramount) AS bigSpend FROM customers AS c FULL JOIN rentals AS r ON c.cid = r.cid GROUP BY c.cid ORDER BY bigSpend DESC;



/* AVG */

SELECT AVG(ramount) FROM rentals;

SELECT c.cid, c.cname, AVG(r.ramount) AS bigSpend FROM customers AS c FULL JOIN rentals AS r ON c.cid = r.cid GROUP BY c.cid;



/* MAX */
SELECT MAX(ramount) FROM rentals;

SELECT c.cid, c.cname, MAX(r.ramount) AS bigSpend FROM customers AS c FULL JOIN rentals AS r ON c.cid = r.cid GROUP BY c.cid;



/* MIN */
SELECT MIN(ramount) FROM rentals;

SELECT c.cid, c.cname, MIN(r.ramount) AS bigSpend FROM customers AS c FULL JOIN rentals AS r ON c.cid = r.cid GROUP BY c.cid;


/* AFTER CLASS */

SELECT cu.cname AS firstName FROM customers AS cu;