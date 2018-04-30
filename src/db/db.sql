CREATE DATABASE webchat DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

use webchat;

CREATE TABLE IF NOT EXISTS `user` (
  `id` integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(18) NOT NULL DEFAULT '',
  `isonline` varchar(10) NOT NULL DEFAULT '',
  `password` varchar(48) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `file` (
  `id` integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `filename` varchar(100) NOT NULL DEFAULT '',
  `filepath` varchar(100) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `message` (
  `id` integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(18) NOT NULL DEFAULT '',
  `time`     varchar(20) NOT NULL DEFAULT '',
  `body`     varchar(500) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8;