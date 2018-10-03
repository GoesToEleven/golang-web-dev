CREATE DATABASE entertainers;

\l

\c entertainers

CREATE TABLE entertainers (eid INT PRIMARY KEY NOT NULL, ename TEXT NOT NULL);

\d

\d entertainers

CREATE TABLE engagements (enID INT PRIMARY KEY NOT NULL, eid INT NOT NULL REFERENCES entertainers(eid), encontractprice INT NOT NULL);

INSERT INTO entertainers VALUES (1, 'WOOKIE-WOOKIE'), (2, 'HodgePodge'), (3, 'ScarlettLetters'), (4, 'TomorrowsSun'), (5, 'WhereWithAll'), (6, 'CowboyFutures'), (7, 'DesperateTortillas'), (8, 'CaldenarFlippers'), (9, 'LaughingJackals'), (10, 'CareeningForthright');


generate values
https://play.golang.org/p/iET0hmFrcgf

INSERT INTO engagements VALUES (1,2,9888),(2,8,6060),(3,2,3319),(4,6,4541),(5,7,5301),(6,5,10512),(7,3,7090),(8,9,5275),(9,2,3446),(10,8,11107),(11,6,7467),(12,9,8259),(13,8,11948),(14,8,4889),(15,1,5016),(16,2,2409),(17,8,8832),(18,10,7357),(19,8,2632),(20,6,7027),(21,4,5091),(22,5,2564),(23,4,6148),(24,9,6325),(25,10,3354),(26,8,5722),(27,10,4200),(28,1,10706),(29,9,6539),(30,4,11356),(31,2,10511),(32,6,2157),(33,7,11829),(34,2,9203),(35,4,7747),(36,4,6377),(37,3,11719),(38,8,7095),(39,8,9464),(40,7,8421),(41,4,2954),(42,8,5134),(43,2,2060),(44,4,10644),(45,2,4003),(46,9,11337),(47,7,11108),(48,1,8504),(49,3,11844),(50,6,3599),(51,6,3352),(52,6,11758),(53,8,10011),(54,1,7286),(55,1,5633),(56,9,10554),(57,2,10583),(58,5,3298),(59,8,8138),(60,2,7895),(61,7,7803),(62,2,4080),(63,7,3271),(64,4,5087),(65,10,10982),(66,3,9176),(67,6,7711),(68,8,5750),(69,9,4819),(70,5,9904),(71,5,6548),(72,3,3533),(73,7,9840),(74,1,7787),(75,2,10077),(76,1,9352),(77,5,2365),(78,6,11184),(79,2,2091),(80,3,4259),(81,8,5232),(82,9,10155),(83,3,3224),(84,3,6209),(85,4,3969),(86,7,5711),(87,6,2441),(88,5,5163),(89,8,6416),(90,2,5040),(91,1,11514),(92,1,3360),(93,1,2784),(94,1,4985),(95,8,10011),(96,6,6163),(97,10,9921),(98,9,8757),(99,6,10667),(100,1,11457);

SELECT e.ename, en.encontractprice FROM entertainers AS e INNER JOIN engagements AS en ON e.eid = en.eid;

SELECT e.ename, en.encontractprice FROM entertainers AS e JOIN engagements AS en ON e.eid = en.eid;

INSERT INTO entertainers VALUES (11, 'SomebodyNobodyLoves');

SELECT e.ename, en.encontractprice FROM entertainers AS e LEFT JOIN engagements AS en ON e.eid = en.eid;

SELECT e.ename, en.encontractprice FROM entertainers AS e FULL JOIN engagements AS en ON e.eid = en.eid;

SELECT COUNT(encontractprice) as TotalEngagements FROM engagements;
SELECT SUM(encontractprice) as TotalRevenue FROM engagements;
SELECT MAX(encontractprice) as MaxContractPrice FROM engagements;
SELECT MIN(encontractprice) as MinContractPrice FROM engagements;
SELECT AVG(encontractprice) as AvgContractPrice FROM engagements;

SELECT e.ename, MAX(en.encontractprice) as MaxContractPrice FROM entertainers AS e JOIN engagements AS en ON e.eid = en.eid GROUP BY e.ename;

SELECT e.ename, MIN(en.encontractprice) as MinContractPrice FROM entertainers AS e JOIN engagements AS en ON e.eid = en.eid GROUP BY e.ename ORDER BY MinContractPrice;

SELECT e.ename, AVG(en.encontractprice) as AvgContractPrice FROM entertainers AS e JOIN engagements AS en ON e.eid = en.eid GROUP BY e.ename ORDER BY AvgContractPrice;

SELECT e.ename, SUM(en.encontractprice) as SumContractPrice FROM entertainers AS e JOIN engagements AS en ON e.eid = en.eid GROUP BY e.ename ORDER BY SumContractPrice;

SELECT e.ename, COUNT(en.encontractprice) as CountContractPrice FROM entertainers AS e JOIN engagements AS en ON e.eid = en.eid GROUP BY e.ename ORDER BY CountContractPrice;

SELECT e.ename, COUNT(en.encontractprice) AS NumContracts, SUM(en.encontractprice) AS TotalRevenue, MIN(en.encontractprice) AS MinRev, MAX(en.encontractprice) AS MaxRev, AVG(en.encontractprice) AS AvgRev FROM entertainers AS e JOIN engagements AS en ON e.eid = en.eid GROUP BY e.ename ORDER BY NumContracts;

SELECT e.ename, COUNT(en.encontractprice) AS NumContracts, SUM(en.encontractprice) AS TotalRevenue, MIN(en.encontractprice) AS MinRev, MAX(en.encontractprice) AS MaxRev, AVG(en.encontractprice) AS AvgRev FROM entertainers AS e LEFT JOIN engagements AS en ON e.eid = en.eid GROUP BY e.ename ORDER BY NumContracts;

SELECT e.ename, COUNT(*) AS NumContracts, SUM(en.encontractprice) AS TotalRevenue, MIN(en.encontractprice) AS MinRev, MAX(en.encontractprice) AS MaxRev, AVG(en.encontractprice) AS AvgRev FROM entertainers AS e LEFT JOIN engagements AS en ON e.eid = en.eid GROUP BY e.ename ORDER BY NumContracts;



/* multiple fields in group by*/

DROP TABLE engagements;

DROP TABLE entertainers;

CREATE TABLE entertainers (eid INT PRIMARY KEY NOT NULL, ename TEXT NOT NULL, ecity TEXT NOT NULL);

INSERT INTO entertainers VALUES (1, 'WOOKIE-WOOKIE', 'Paris'), (2, 'HodgePodge', 'Paris'), (3, 'ScarlettLetters', 'Paris'), (4, 'TomorrowsSun', 'Hollywood'), (5, 'WhereWithAll', 'Hollywood'), (6, 'CowboyFutures', 'Hollywood'), (7, 'DesperateTortillas', 'New York'), (8, 'CaldenarFlippers', 'New York'), (9, 'LaughingJackals', 'San Francisco'), (10, 'CareeningForthright', 'San Francisco');

CREATE TABLE engagements (enID INT PRIMARY KEY NOT NULL, eid INT NOT NULL REFERENCES entertainers(eid), encontractprice INT NOT NULL);

generate values
https://play.golang.org/p/iET0hmFrcgf

INSERT INTO engagements VALUES (1,2,9888),(2,8,6060),(3,2,3319),(4,6,4541),(5,7,5301),(6,5,10512),(7,3,7090),(8,9,5275),(9,2,3446),(10,8,11107),(11,6,7467),(12,9,8259),(13,8,11948),(14,8,4889),(15,1,5016),(16,2,2409),(17,8,8832),(18,10,7357),(19,8,2632),(20,6,7027),(21,4,5091),(22,5,2564),(23,4,6148),(24,9,6325),(25,10,3354),(26,8,5722),(27,10,4200),(28,1,10706),(29,9,6539),(30,4,11356),(31,2,10511),(32,6,2157),(33,7,11829),(34,2,9203),(35,4,7747),(36,4,6377),(37,3,11719),(38,8,7095),(39,8,9464),(40,7,8421),(41,4,2954),(42,8,5134),(43,2,2060),(44,4,10644),(45,2,4003),(46,9,11337),(47,7,11108),(48,1,8504),(49,3,11844),(50,6,3599),(51,6,3352),(52,6,11758),(53,8,10011),(54,1,7286),(55,1,5633),(56,9,10554),(57,2,10583),(58,5,3298),(59,8,8138),(60,2,7895),(61,7,7803),(62,2,4080),(63,7,3271),(64,4,5087),(65,10,10982),(66,3,9176),(67,6,7711),(68,8,5750),(69,9,4819),(70,5,9904),(71,5,6548),(72,3,3533),(73,7,9840),(74,1,7787),(75,2,10077),(76,1,9352),(77,5,2365),(78,6,11184),(79,2,2091),(80,3,4259),(81,8,5232),(82,9,10155),(83,3,3224),(84,3,6209),(85,4,3969),(86,7,5711),(87,6,2441),(88,5,5163),(89,8,6416),(90,2,5040),(91,1,11514),(92,1,3360),(93,1,2784),(94,1,4985),(95,8,10011),(96,6,6163),(97,10,9921),(98,9,8757),(99,6,10667),(100,1,11457);

SELECT e.ecity, e.ename, COUNT(*) AS NumContracts, SUM(en.encontractprice) AS TotalRevenue, MIN(en.encontractprice) AS MinRev, MAX(en.encontractprice) AS MaxRev, AVG(en.encontractprice) AS AvgRev FROM entertainers AS e LEFT JOIN engagements AS en ON e.eid = en.eid GROUP BY e.ecity, e.ename ORDER BY e.ecity, NumContracts;

/* simulating DISTINCT */

SELECT ecity FROM entertainers;

SELECT DISTINCT(ecity) FROM entertainers;

SELECT ecity FROM entertainers GROUP BY ecity;

SELECT ecity, COUNT(*) AS bandspercity FROM entertainers GROUP BY ecity;


