-- CREATE DATABASE gostudydb CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE gostudydb;

DROP TABLE IF EXISTS `do_something`;

CREATE TABLE `do_something`
(
    `id`                   bigint(20) NOT NULL AUTO_INCREMENT,
    `flag`                 tinyint(1) NOT NULL DEFAULT '0',
    `create_time_utc`      datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `last_update_time_utc` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `who`                  varchar(250) NOT NULL,
    `action`               varchar(250) NOT NULL,
    `thing`                varchar(250) NOT NULL,
    `more`                 varchar(250) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uidx_do_something_01` (`who`, `action`, `thing`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='do something';