-- MySQL dump 10.16  Distrib 10.1.43-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: cs340
-- ------------------------------------------------------
-- Server version	10.1.43-MariaDB-0ubuntu0.18.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Temporary table structure for view `ALL_STICKY_POSTS`
--

DROP TABLE IF EXISTS `ALL_STICKY_POSTS`;
/*!50001 DROP VIEW IF EXISTS `ALL_STICKY_POSTS`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE TABLE `ALL_STICKY_POSTS` (
  `pID` tinyint NOT NULL,
  `Title` tinyint NOT NULL,
  `Timestamp` tinyint NOT NULL,
  `Summary` tinyint NOT NULL,
  `Body` tinyint NOT NULL,
  `uID` tinyint NOT NULL
) ENGINE=MyISAM */;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `Comment`
--

DROP TABLE IF EXISTS `Comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Comment` (
  `cID` int(11) NOT NULL AUTO_INCREMENT,
  `timestamp` datetime NOT NULL,
  `content` varchar(256) NOT NULL,
  `uID` int(11) NOT NULL,
  `pID` int(11) NOT NULL,
  PRIMARY KEY (`cID`),
  KEY `uID` (`uID`),
  KEY `Comment_ibfk_2` (`pID`),
  CONSTRAINT `Comment_ibfk_1` FOREIGN KEY (`uID`) REFERENCES `User` (`uID`),
  CONSTRAINT `Comment_ibfk_2` FOREIGN KEY (`pID`) REFERENCES `Post` (`pID`)
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Comment`
--

LOCK TABLES `Comment` WRITE;
/*!40000 ALTER TABLE `Comment` DISABLE KEYS */;
INSERT INTO `Comment` VALUES (66,'2019-12-12 08:18:24','test comment',12,45),(67,'2019-12-12 08:34:31','This is a comment on a post, try deleting me!',12,46);
/*!40000 ALTER TABLE `Comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Post`
--

DROP TABLE IF EXISTS `Post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Post` (
  `pID` int(11) NOT NULL AUTO_INCREMENT,
  `timestamp` datetime NOT NULL,
  `title` varchar(100) NOT NULL,
  `summary` varchar(200) NOT NULL,
  `body` varchar(2048) NOT NULL,
  `uID` int(11) NOT NULL,
  PRIMARY KEY (`pID`),
  KEY `uID` (`uID`),
  CONSTRAINT `Post_ibfk_1` FOREIGN KEY (`uID`) REFERENCES `User` (`uID`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Post`
--

LOCK TABLES `Post` WRITE;
/*!40000 ALTER TABLE `Post` DISABLE KEYS */;
INSERT INTO `Post` VALUES (41,'2019-12-12 07:31:54','test','test','test',12),(45,'2019-12-12 08:14:54','Thanks for a great term!','Agenda for Winter break.','Meetings will start next week, and so I will post a time before Sunday as to when we will meet during winter break. The main goals of Winter break are as follows:\n- Have an exact blueprint of what needs to be modeled in CAD. (The design implemented from various papers should be analyzed and combined into one perfect design we are going to use.)\n- AI should be ready for collected data, so that we can plug in data easily, and output a neural network that can predict standing from walking.\n- Data collection method should be thoroughly specified and hopefully implemented (although we might need to start this Winter Term)\n- Plan ORGANIZED presentations for Week 1 of Winter Term.\n- Plan the Event happening Week 5 of Winter Term.',13),(46,'2019-12-12 08:34:08','Update Post','...','This is an example of an update post',12),(47,'2019-12-12 09:00:35','Unverified post','I\'m unverified!','You should try verifying me!',13);
/*!40000 ALTER TABLE `Post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Post_Tags`
--

DROP TABLE IF EXISTS `Post_Tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Post_Tags` (
  `pID` int(11) NOT NULL,
  `tag` varchar(16) NOT NULL,
  PRIMARY KEY (`pID`,`tag`),
  KEY `pID` (`pID`),
  KEY `pID_2` (`pID`),
  CONSTRAINT `pID` FOREIGN KEY (`pID`) REFERENCES `Post` (`pID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Post_Tags`
--

LOCK TABLES `Post_Tags` WRITE;
/*!40000 ALTER TABLE `Post_Tags` DISABLE KEYS */;
INSERT INTO `Post_Tags` VALUES (41,'verified'),(45,'announcement'),(45,'sticky'),(45,'verified'),(46,'update'),(46,'verified');
/*!40000 ALTER TABLE `Post_Tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Stickied_Posts`
--

DROP TABLE IF EXISTS `Stickied_Posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Stickied_Posts` (
  `pID` int(11) NOT NULL,
  KEY `Stickied_Posts_ibfk_1` (`pID`),
  CONSTRAINT `Stickied_Posts_ibfk_1` FOREIGN KEY (`pID`) REFERENCES `Post` (`pID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Stickied_Posts`
--

LOCK TABLES `Stickied_Posts` WRITE;
/*!40000 ALTER TABLE `Stickied_Posts` DISABLE KEYS */;
INSERT INTO `Stickied_Posts` VALUES (45);
/*!40000 ALTER TABLE `Stickied_Posts` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`admin`@`localhost`*/ /*!50003 TRIGGER `checkNumberOfSPost` BEFORE INSERT ON `Stickied_Posts` FOR EACH ROW BEGIN
    IF (SELECT COUNT(*) FROM Stickied_Posts) = 5 THEN
        DELETE FROM Stickied_Posts
        ORDER BY pID 
        LIMIT 1;
    END IF;    
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `User` (
  `isAdmin` tinyint(1) NOT NULL DEFAULT '0',
  `uID` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(20) NOT NULL,
  `email` varchar(128) DEFAULT NULL,
  `saltedhash` varchar(256) NOT NULL,
  `role` varchar(16) NOT NULL DEFAULT 'Member',
  PRIMARY KEY (`uID`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES (0,7,'John','john@yourteam.solutions','$2a$04$/QJyTzfxY3yb.hGvsSMpX.S99aWQ9P8ZRDPPusJNv7YSvehyYynGK','Member'),(0,10,'test','test@test.test','$2a$04$1LKVnSDq.woeXxgwssb1wut0mE5Tak3/gjboOpacX862ddqVGeMeW','Member'),(1,12,'John Warila','johnwarila@gmail.com','$2a$04$n2HVTO3UQ0YwiX4CYk5S8us0BlRNirQFdSD4p.buUYs0jsTDqb1g.','Member'),(1,13,'Srikar Valluri','srikar@yourteam.solutions','$2a$04$rbMBAZh.WBffKSFfBT/yg.YhyCBZQZnAeJSX26O8z/1cLUQOItc0i','Member'),(1,14,'Admin Grader','admin@tdc.com','$2a$04$yUIV3Z1prhr3yqea8gCXReRKD8gXsBOSGKdkMQsM/m/6T.tVnmwjK','Member');
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Final view structure for view `ALL_STICKY_POSTS`
--

/*!50001 DROP TABLE IF EXISTS `ALL_STICKY_POSTS`*/;
/*!50001 DROP VIEW IF EXISTS `ALL_STICKY_POSTS`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_unicode_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`admin`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `ALL_STICKY_POSTS` AS select `Post`.`pID` AS `pID`,`Post`.`title` AS `Title`,`Post`.`timestamp` AS `Timestamp`,`Post`.`summary` AS `Summary`,`Post`.`body` AS `Body`,`Post`.`uID` AS `uID` from (`Post` join `Stickied_Posts`) where (`Stickied_Posts`.`pID` = `Post`.`pID`) order by `Post`.`timestamp` desc */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-12-12 17:31:00
