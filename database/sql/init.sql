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
INSERT INTO `spots` VALUES ('9fdf8594', '特別史跡五稜郭跡', 41.7969245, 140.7545951);
INSERT INTO `emotions` VALUES ('9fdf8594', 0.2, 0.3, 0.2, 0.0);

INSERT INTO `spots` VALUES ('9fdf8c2e', '五稜郭タワー', 41.7947351, 140.7518855);
INSERT INTO `emotions` VALUES ('9fdf8c2e', 0.0, 0.8, 0.4, 0.3);

INSERT INTO `spots` VALUES ('9fdf8c4c', '旧函館区公会堂', 41.7650242, 140.7067455);
INSERT INTO `emotions` VALUES ('9fdf8c4c', 0.4, 0.2, 0.3, 0.1);

INSERT INTO `spots` VALUES ('9fdf8ca6', '函館ハリストス正教会', 41.7628007, 140.7099986);
INSERT INTO `emotions` VALUES ('9fdf8ca6', 0.7, 0.9, 0.7, 0.5);

INSERT INTO `spots` VALUES ('9fdf8cc4', 'カトリック函館元町教会', 41.763521, 140.7107733);
INSERT INTO `emotions` VALUES ('9fdf8cc4', 0.4, 0.2, 0.6, 0.6);

INSERT INTO `spots` VALUES ('9fdf8ce2', '函館奉行所跡', 41.7968697, 140.7546177);
INSERT INTO `emotions` VALUES ('9fdf8ce2', 0.8, 0.0, 0.8, 0.4);

INSERT INTO `spots` VALUES ('9fdf8d00', '函館市旧イギリス領事館', 41.7659525, 140.708462);
INSERT INTO `emotions` VALUES ('9fdf8d00', 0.4, 0.9, 1.0, 0.7);

INSERT INTO `spots` VALUES ('9fdf8d1e', '函館香雪園', 41.7919847, 140.8006364);
INSERT INTO `emotions` VALUES ('9fdf8d1e', 0.5, 0.7, 0.7, 0.7);

INSERT INTO `spots` VALUES ('9fdf8d3c', '元町公園', 41.7654424, 140.7072904);
INSERT INTO `emotions` VALUES ('9fdf8d3c', 0.3, 0.6, 0.4, 0.4);

INSERT INTO `spots` VALUES ('9fdf8d5a', '湯倉神社', 41.7823088, 140.7887117);
INSERT INTO `emotions` VALUES ('9fdf8d5a', 0.0, 0.7, 0.4, 0.5);

INSERT INTO `spots` VALUES ('9fdf8d78', '函館山展望台', 41.7593546, 140.7025134);
INSERT INTO `emotions` VALUES ('9fdf8d78', 0.5, 1.0, 0.3, 0.4);

INSERT INTO `spots` VALUES ('9fdf8d96', '函館八幡宮', 41.7538784, 140.7078736);
INSERT INTO `emotions` VALUES ('9fdf8d96', 0.1, 0.2, 0.4, 0.2);

INSERT INTO `spots` VALUES ('9fdf8db4', '函館聖ヨハネ教会', 41.7625283, 140.7105957);
INSERT INTO `emotions` VALUES ('9fdf8db4', 0.7, 0.3, 0.8, 0.6);

INSERT INTO `spots` VALUES ('9fdf8dd2', '道南四季の杜公園', 41.8478113, 140.7664765);
INSERT INTO `emotions` VALUES ('9fdf8dd2', 0.4, 0.3, 0.7, 0.9);

INSERT INTO `spots` VALUES ('9fdf8e18', '啄木小公園', 41.7725311, 140.7520065);
INSERT INTO `emotions` VALUES ('9fdf8e18', 0.2, 0.6, 0.6, 0.3);

INSERT INTO `spots` VALUES ('9fdf8e36', '四稜郭跡', 41.7650242, 140.7067455);
INSERT INTO `emotions` VALUES ('9fdf8e36', 0.2, 0.5, 0.5, 0.6);

INSERT INTO `spots` VALUES ('9fdf8e54', '船魂神社', 41.7634255, 140.7078321);
INSERT INTO `emotions` VALUES ('9fdf8e54', 0.0, 0.3, 0.2, 0.6);

INSERT INTO `spots` VALUES ('9fdf8e72', '志苔館跡', 41.7657778, 140.8199579);
INSERT INTO `emotions` VALUES ('9fdf8e72', 0.7, 0.5, 0.6, 0.3);

INSERT INTO `spots` VALUES ('9fdf8e90', '恵山つつじ公園', 41.7924851, 141.1519726);
INSERT INTO `emotions` VALUES ('9fdf8e90', 0.0, 0.0, 0.0, 0.0);

INSERT INTO `spots` VALUES ('9fdf8eae', '旧相馬邸', 41.765246, 140.7087094);
INSERT INTO `emotions` VALUES ('9fdf8eae', 0.4, 0.5, 0.7, 0.3);

INSERT INTO `spots` VALUES ('9fdf8ecc', '函館護国神社', 41.759291, 140.7111636);
INSERT INTO `emotions` VALUES ('9fdf8ecc', 0.1, 0.3, 0.1, 1.0);

INSERT INTO `spots` VALUES ('9fdf8eea', '函館市市民の森', 41.7872573, 140.8176493);
INSERT INTO `emotions` VALUES ('9fdf8eea', 0.3, 0.3, 0.9, 0.4);

INSERT INTO `spots` VALUES ('9fdf8f08', '松前藩戸切地陣屋跡', 41.8540736, 140.6172477);
INSERT INTO `emotions` VALUES ('9fdf8f08', 0.2, 0.3, 0.6, 0.9);

INSERT INTO `spots` VALUES ('9fdf8f26', '八番坂', 41.7680351, 140.6949824);
INSERT INTO `emotions` VALUES ('9fdf8f26', 1.0, 0.1, 0.7, 0.1);

INSERT INTO `spots` VALUES ('9fdf8f44', '城岱牧場展望台', 41.9229559, 140.7088118);
INSERT INTO `emotions` VALUES ('9fdf8f44', 0.8, 0.9, 0.1, 0.5);

INSERT INTO `spots` VALUES ('9fdf8f62', '函館市昭和公園', 41.8148108, 140.7243656);
INSERT INTO `emotions` VALUES ('9fdf8f62', 0.6, 0.1, 0.7, 0.1);

INSERT INTO `spots` VALUES ('9fdf8f80', '千畳敷', 41.7471272, 140.69848);
INSERT INTO `emotions` VALUES ('9fdf8f80', 0.2, 1.0, 0.5, 0.9);

INSERT INTO `spots` VALUES ('9fdf8f9e', '亀田八幡宮', 41.7915686, 140.7343907);
INSERT INTO `emotions` VALUES ('9fdf8f9e', 0.7, 0.9, 0.4, 0.3);

INSERT INTO `spots` VALUES ('9fdf8fbc', '七重浜海浜公園', 41.8139114, 140.702146);
INSERT INTO `emotions` VALUES ('9fdf8fbc', 0.2, 0.1, 1.0, 0.2);

INSERT INTO `spots` VALUES ('9fdf8fda', '入船町前浜海水浴場', 41.7334972, 140.6288006);
INSERT INTO `emotions` VALUES ('9fdf8fda', 0.3, 0.3, 0.6, 0.7);
