CREATE DATABASE music;

\c music

CREATE TABLE musicians (mid INT PRIMARY KEY NOT NULL, mname TEXT NOT NULL);

INSERT INTO musicians VALUES (1, 'Dave Mustang'), (2, 'Steven Tyler'), (3, 'Sam Smith'), (4, 'Sarah McLachlan'), (5, 'Bono'), (6, 'Steve Miller'), (7, 'Bruce Springsteen');

CREATE TABLE bands (bid INT PRIMARY KEY NOT NULL, bname TEXT NOT NULL);

INSERT INTO bands VALUES (1, 'Aerosmith'), (2, 'U2'), (3, 'Ghost'), (4, 'E Street'), (5, 'Metallica'), (6, 'Megadeath');

CREATE TABLE bandmembers (bmid INT PRIMARY KEY NOT NULL, bid INT REFERENCES bands(bid) NOT NULL, mid INT REFERENCES musicians(mid) NOT NULL);

INSERT INTO bandmembers VALUES (1,1,1), (2,2,2), (3,3,3), (4,4,4), (5,5,5), (6,5,7), (7,4,4), (8,3,3);

CREATE TABLE styles (sid INT PRIMARY KEY NOT NULL, sname TEXT NOT NULL);

INSERT INTO styles VALUES (1,'jazz'), (2,'metal'), (3,'rock'), (4,'country'), (5,'pop'), (6,'folk');

CREATE TABLE bandstyles (bsid INT PRIMARY KEY NOT NULL, bid INT REFERENCES bands(bid) NOT NULL, sid INT REFERENCES styles(sid) NOT NULL);

INSERT INTO bandstyles VALUES (1,1,1), (2,2,2), (3,3,3), (4,4,4), (5,5,5), (6,4,4), (7,3,3);

------------ queries ---------------


SELECT m.mname, b.bname, s.sname FROM (((musicians AS m FULL JOIN bandmembers AS bm ON m.mid = bm.mid) FULL JOIN bands AS b ON bm.bid = b.bid) FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid;


/* generate more band members */
https://play.golang.org/p/uv2g4Ap21Sl

INSERT INTO bandmembers VALUES (9,2,3),(10,3,5),(11,2,4),(12,1,1),(13,2,1),(14,5,2),(15,3,5),(16,4,5),(17,2,1),(18,3,2),(19,1,2),(20,4,4),(21,3,3),(22,3,4),(23,1,1),(24,2,4),(25,3,2),(26,5,2),(27,3,2),(28,1,2),(29,4,1),(30,5,4),(31,4,3),(32,4,5),(33,5,4),(34,3,2),(35,5,5),(36,1,1),(37,4,4),(38,4,1),(39,2,1),(40,1,2);


SELECT bname, sname FROM (bands AS b LEFT JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid;

SELECT m.mname, b.bname FROM (musicians AS m LEFT JOIN bandmembers AS bm ON m.mid = bm.mid) LEFT JOIN bands AS b ON bm.bid = b.bid;

SELECT m.mname, b.bname FROM (musicians AS m LEFT JOIN bandmembers AS bm ON m.mid = bm.mid) LEFT JOIN bands AS b ON bm.bid = b.bid GROUP BY m.mname, b.bname;

SELECT m.mid, m.mname, b.bname FROM (musicians AS m LEFT JOIN bandmembers AS bm ON m.mid = bm.mid) LEFT JOIN bands AS b ON bm.bid = b.bid GROUP BY m.mid, m.mname, b.bname;

/* does not work*/
SELECT b.bname, s.sname, COUNT(*) AS members FROM (((musicians AS m JOIN bandmembers AS bm ON m.mid = bm.mid) JOIN bands AS b ON bm.bid = b.bid) JOIN bandstyles AS bs ON b.bid = bs.bid) JOIN styles AS s ON bs.sid = s.sid;

SELECT COUNT(*) AS total FROM musicians;

/* does not work*/
SELECT COUNT(*) AS total FROM musicians GROUP BY mname;

SELECT mname, COUNT(*) AS total FROM musicians GROUP BY mname;

SELECT b.bname, s.sname, COUNT(*) AS members FROM (((musicians AS m JOIN bandmembers AS bm ON m.mid = bm.mid) JOIN bands AS b ON bm.bid = b.bid) JOIN bandstyles AS bs ON b.bid = bs.bid) JOIN styles AS s ON bs.sid = s.sid GROUP BY b.bname, s.sname;

SELECT b.bname, s.sname, COUNT(*) AS members FROM (((musicians AS m JOIN bandmembers AS bm ON m.mid = bm.mid) JOIN bands AS b ON bm.bid = b.bid) JOIN bandstyles AS bs ON b.bid = bs.bid) JOIN styles AS s ON bs.sid = s.sid GROUP BY b.bname, s.sname HAVING COUNT(*) > 10;

SELECT b.bname, s.sname, COUNT(bm.bmid) AS bandmembers FROM (((musicians AS m JOIN bandmembers AS bm ON m.mid = bm.mid) JOIN bands AS b ON bm.bid = b.bid) JOIN bandstyles AS bs ON b.bid = bs.bid) JOIN styles AS s ON bs.sid = s.sid GROUP BY b.bname, s.sname;

SELECT b.bname, s.sname, COUNT(bm.bmid) AS bandmembers FROM (((musicians AS m JOIN bandmembers AS bm ON m.mid = bm.mid) JOIN bands AS b ON bm.bid = b.bid) JOIN bandstyles AS bs ON b.bid = bs.bid) JOIN styles AS s ON bs.sid = s.sid GROUP BY b.bname, s.sname HAVING COUNT(*) > 10;

