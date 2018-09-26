CREATE DATABASE hollywood;

\l

\c hollywood

CREATE TABLE customers (cid INT PRIMARY KEY NOT NULL, cname TEXT);

INSERT INTO customers VALUES (1,'Charlie'), (2,'Chris'), (3,'Christina'), (4,'Cassandra'), (5,'Cull'), (6,'Cabron'), (7,'Carla'), (8,'Caitland'), (9,'Colleen'), (10,'Christian');

CREATE TABLE movies (mid INT PRIMARY KEY NOT NULL, mname TEXT);

INSERT INTO movies VALUES (1,'Jaws'), (2,'First Blood'), (3,'Rambo'), (4,'Terms of Endearment'), (5,'Heat'), (6,'The Lives of Others'), (7,'A Man Called Ove'), (8,'Disconnect'), (9,'Ex Machina'), (10,'Castle');

CREATE TABLE actors (aid INT PRIMARY KEY NOT NULL, aname TEXT);

INSERT INTO actors VALUES (1,'Alfred'), (2,'Albert'), (3,'Angel'), (4,'Angelina'), (5,'Angie'), (6,'Agnes'), (7,'Amnia'), (8,'Allison'), (9,'Alejandro'), (10,'Aerial');

CREATE TABLE rentals (rid INT PRIMARY KEY NOT NULL, cid INT NOT NULL REFERENCES customers(cid), mid INT NOT NULL REFERENCES movies(mid));

INSERT INTO rentals VALUES (1,2,8),(2,8,10),(3,2,9),(4,6,1),(5,7,1),(6,5,2),(7,3,10),(8,9,5),(9,2,6),(10,8,7),(11,6,7),(12,9,9),(13,8,8),(14,8,9),(15,1,6),(16,2,9),(17,8,2),(18,10,7),(19,8,2),(20,6,7),(21,4,1),(22,5,4),(23,4,8),(24,9,5),(25,10,4),(26,8,2),(27,10,10),(28,1,6),(29,9,9),(30,4,6),(31,2,1),(32,6,7),(33,7,9),(34,2,3),(35,4,7),(36,4,7),(37,3,9),(38,8,5),(39,8,4),(40,7,1),(41,4,4),(42,8,4),(43,2,10),(44,4,4),(45,2,3),(46,9,7),(47,7,8),(48,1,4),(49,3,4),(50,6,9),(51,6,2),(52,6,8),(53,8,1),(54,1,6),(55,1,3),(56,9,4),(57,2,3),(58,5,8),(59,8,8),(60,2,5),(61,7,3),(62,2,10),(63,7,1),(64,4,7),(65,10,2),(66,3,6),(67,6,1),(68,8,10),(69,9,9),(70,5,4),(71,5,8),(72,3,3),(73,7,10),(74,1,7),(75,2,7),(76,1,2),(77,5,5),(78,6,4),(79,2,1),(80,3,9),(81,8,2),(82,9,5),(83,3,4),(84,3,9),(85,4,9),(86,7,1),(87,6,1),(88,5,3),(89,8,6),(90,2,10),(91,1,4),(92,1,10),(93,1,4),(94,1,5),(95,8,1),(96,6,3),(97,10,1),(98,9,7),(99,6,7),(100,1,7);

CREATE TABLE castmembers (caid INT PRIMARY KEY NOT NULL, mid INT NOT NULL REFERENCES movies(mid), aid INT NOT NULL REFERENCES actors(aid));

INSERT INTO castmembers VALUES (1,2,8),(2,8,10),(3,2,9),(4,6,1),(5,7,1),(6,5,2),(7,3,10),(8,9,5),(9,2,6),(10,8,7),(11,6,7),(12,9,9),(13,8,8),(14,8,9),(15,1,6),(16,2,9),(17,8,2),(18,10,7),(19,8,2),(20,6,7),(21,4,1),(22,5,4),(23,4,8),(24,9,5),(25,10,4),(26,8,2),(27,10,10),(28,1,6),(29,9,9),(30,4,6),(31,2,1),(32,6,7),(33,7,9),(34,2,3),(35,4,7),(36,4,7),(37,3,9),(38,8,5),(39,8,4),(40,7,1),(41,4,4),(42,8,4),(43,2,10),(44,4,4),(45,2,3),(46,9,7),(47,7,8),(48,1,4),(49,3,4),(50,6,9),(51,6,2),(52,6,8),(53,8,1),(54,1,6),(55,1,3),(56,9,4),(57,2,3),(58,5,8),(59,8,8),(60,2,5),(61,7,3),(62,2,10),(63,7,1),(64,4,7),(65,10,2),(66,3,6),(67,6,1),(68,8,10),(69,9,9),(70,5,4),(71,5,8),(72,3,3),(73,7,10),(74,1,7),(75,2,7),(76,1,2),(77,5,5),(78,6,4),(79,2,1),(80,3,9),(81,8,2),(82,9,5),(83,3,4),(84,3,9),(85,4,9),(86,7,1),(87,6,1),(88,5,3),(89,8,6),(90,2,10),(91,1,4),(92,1,10),(93,1,4),(94,1,5),(95,8,1),(96,6,3),(97,10,1),(98,9,7),(99,6,7),(100,1,7);


SELECT c.cname, m.mname FROM customers AS c JOIN rentals AS r ON c.cid = r.cid JOIN movies AS m ON r.mid = m.mid;

SELECT a.aname, m.mname FROM actors AS a JOIN castmembers AS cm ON  a.aid = cm.aid JOIN movies AS m ON cm.mid = m.mid;

INSERT INTO customers VALUES (101, 'Mike'), (102, 'Max');

INSERT INTO movies VALUES (101, 'Aliens'), (102, 'Ragoon');

INSERT INTO actors VALUES (101, 'Milfred'), (102, 'Martine');

SELECT c.cname, m.mname, a.aname FROM customers AS c FULL JOIN rentals AS r ON c.cid = r.cid FULL JOIN movies AS m ON r.mid = m.mid FULL JOIN castmembers AS cm ON m.mid = cm.mid FULL JOIN actors AS a ON cm.aid = a.aid;

