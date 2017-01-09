DROP DATABASE IF EXISTS `app`;
CREATE DATABASE /*!32312 IF NOT EXISTS*/ `app` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `app`;

CREATE TABLE IF NOT EXISTS `todos` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `text` varchar(200) DEFAULT NULL,
  `checked` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_todos_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

INSERT INTO `todos` (`id`,`created_at`,`updated_at`,`deleted_at`,`text`,`checked`) VALUES (1,'2017-01-09 10:58:53','2017-01-09 10:58:53',NULL,'Hi',0);
