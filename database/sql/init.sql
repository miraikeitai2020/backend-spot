DROP DATABASE IF EXISTS michishirube;
CREATE DATABASE michishirube;
USE michishirube;

---- create user ----
CREATE USER IF NOT EXISTS 'miraiketai2020'@'%' IDENTIFIED BY 'miraiketai2020';
GRANT ALL ON michishirube.* TO miraiketai2020;

---- drop ----
DROP TABLE IF EXISTS `spots`;
DROP TABLE IF EXISTS `emotions`;
DROP TABLE IF EXISTS `detours`;
DROP TABLE IF EXISTS `images`;

---- create ----
CREATE TABLE IF NOT EXISTS `spots`
(
 `id`               VARCHAR(256),
 `name`             VARCHAR(256) NOT NULL,
 `latitude`         DOUBLE NOT NULL,
 `longitude`        DOUBLE NOT NULL,
 PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `emotions`
(
 `id`               VARCHAR(256),
 `happiness`        DOUBLE NOT NULL,
 `natural`          DOUBLE NOT NULL,
 `sadness`          DOUBLE NOT NULL,
 `anger`            DOUBLE NOT NULL,
 PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `detours`
(
 `id`               VARCHAR(256),
 `name`             VARCHAR(256) NOT NULL,
 `description`      VARCHAR(256) NOT NULL,
 `latitude`         DOUBLE NOT NULL,
 `longitude`        DOUBLE NOT NULL,
 PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS `images`
(
 `id`              VARCHAR(256),
 `image`           LONGBLOB NOT NULL,
 PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

---- insert ----
