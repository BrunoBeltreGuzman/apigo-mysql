DROP TABLE IF EXISTS sessionsfull;

CREATE TABLE `sessionsfull` (
       `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'primary key',
       `user` int DEFAULT NULL,
       `token` varchar(6000) DEFAULT NULL,
       `ip` varchar(30) DEFAULT NULL,
       `createAt` datetime DEFAULT now(),
       `expireAt` datetime DEFAULT now()
) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8;

INSERT INTO
       sessionsfull (user, token, ip, expireAt)
VALUES
       (
              1,
              "as2d31as354f65saf4",
              "213.214.152.1",
              DATE_ADD(now(), INTERVAL 15 MINUTE)
       );

SELECT
       *
FROM
       sessionsfull;

SELECT
       sessionsfull.user,
       sessionsfull.token,
       sessionsfull.ip,
       sessionsfull.expireAt
FROM
       sessionsfull
WHERE
       expireAt > now();

SELECT
       sessionsfull.user,
       sessionsfull.token,
       sessionsfull.ip,
       sessionsfull.expireAt
FROM
       sessionsfull
WHERE
       expireAt > now()
       AND token = 'as2d31as354f65saf4';