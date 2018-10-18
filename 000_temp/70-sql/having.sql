SELECT m.mname, b.bname, s.sname FROM (((musicians AS m FULL JOIN bandmembers AS bm ON m.mid = bm.mid) FULL JOIN bands AS b ON bm.bid = b.bid) FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid;

SELECT s.sname, COUNT(b.bname) AS bands FROM (((musicians AS m FULL JOIN bandmembers AS bm ON m.mid = bm.mid) FULL JOIN bands AS b ON bm.bid = b.bid) FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid GROUP BY s.sname;

SELECT s.sname, COUNT(DISTINCT(b.bname)) AS bands FROM (((musicians AS m FULL JOIN bandmembers AS bm ON m.mid = bm.mid) FULL JOIN bands AS b ON bm.bid = b.bid) FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid GROUP BY s.sname;

SELECT DISTINCT(b.bname) AS name FROM bands AS b;

SELECT DISTINCT(bname) FROM bands;

SELECT * FROM bands;

INSERT INTO bands VALUES (7, 'Madonna');
INSERT INTO bands VALUES (8, 'Madonna');

SELECT b.bname, s.sname FROM (bands AS b FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid;

INSERT INTO bandstyles VALUES (8,5,2), (9,4,1), (10,3,2), (11,2,3), (12,1,4), (13,1,5), (14,1,6);

SELECT m.mname, b.bname, s.sname FROM (((musicians AS m FULL JOIN bandmembers AS bm ON m.mid = bm.mid) FULL JOIN bands AS b ON bm.bid = b.bid) FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid;

SELECT s.sname, COUNT(b.bname) AS bands FROM (bands AS b FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid GROUP BY s.sname;

SELECT s.sname, COUNT(m.mname) AS musicians FROM (((musicians AS m FULL JOIN bandmembers AS bm ON m.mid = bm.mid) FULL JOIN bands AS b ON bm.bid = b.bid) FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid GROUP BY s.sname;

SELECT s.sname, COUNT(DISTINCT(b.bname)) AS bands FROM (bands AS b FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid GROUP BY s.sname;

SELECT bname, sname FROM (bands AS b FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid;

SELECT sname, bname FROM (bands AS b FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid ORDER BY bname;

SELECT DISTINCT bname, sname AS bands FROM (bands AS b FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid ORDER BY bname;

SELECT DISTINCT sname, bname AS bands FROM (bands AS b FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid ORDER BY bname;

SELECT DISTINCT sname, COUNT(bname) AS bands FROM (bands AS b FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid GROUP BY sname;

SELECT sname, COUNT(bname) AS bands FROM (bands AS b FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid GROUP BY sname;

SELECT DISTINCT sname, bname FROM (bands AS b FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid;

SELECT sname, bname FROM (bands AS b FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid;

SELECT sname, COUNT(bname) as bands FROM (bands AS b FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid GROUP BY sname;

SELECT DISTINCT sname, COUNT(bname) as bands FROM (bands AS b FULL JOIN bandstyles AS bs ON b.bid = bs.bid) FULL JOIN styles AS s ON bs.sid = s.sid GROUP BY sname;


















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

