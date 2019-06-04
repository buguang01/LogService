/*
SQLyog Ultimate v12.08 (32 bit)
MySQL - 5.6.43-log : Database - gamelog_db
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


/*Table structure for table `loginfo` */

CREATE TABLE `loginfo` (
  `uid` bigint(20) NOT NULL AUTO_INCREMENT,
  `topicid1` int(11) NOT NULL,
  `topicid2` int(11) NOT NULL,
  `topicid3` int(11) NOT NULL,
  `memberid` int(11) NOT NULL,
  `serviceid` int(11) NOT NULL,
  `uptime` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
  `upnum` bigint(20) NOT NULL,
  `total` bigint(20) NOT NULL,
  `datas` varchar(4000) NOT NULL,
  PRIMARY KEY (`uid`),
  KEY `topicid1` (`memberid`,`topicid1`,`topicid2`,`topicid3`),
  KEY `uptime` (`uptime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Table structure for table `logtree` */

CREATE TABLE `logtree` (
  `topicid1` int(11) NOT NULL,
  `topicid2` int(11) NOT NULL,
  `topicid3` int(11) NOT NULL,
  PRIMARY KEY (`topicid1`,`topicid2`,`topicid3`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Table structure for table `topicinfo` */

CREATE TABLE `topicinfo` (
  `uid` int(11) NOT NULL,
  `name` varchar(128) NOT NULL,
  PRIMARY KEY (`uid`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
