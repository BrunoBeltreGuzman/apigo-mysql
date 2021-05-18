SELECT
       now();

CREATE TABLE mytime (
       id int NOT NULL primary key AUTO_INCREMENT comment 'primary key',
       time1 time,
       time2 time,
       time3 time
) default charset utf8 comment ''
INSERT into
       mytime (time1, time2, time3)
VAlUES
       ('16:53:24', '16:53:24', '16:53:24');

INSERT into
       mytime (time1, time2, time3)
VAlUES
       (now(), '16:53:24', '16:53:24');

SELECT
       mytime.time1,
       mytime.time2
FROM
       mytime
where
       mytime.time1 = mytime.time2
       and mytime.id = 1;

SELECT
       mytime.time1,
       mytime.time2
FROM
       mytime
where
       mytime.time1 = mytime.time2
       and mytime.id = 2;

SELECT
       mytime.time1,
       mytime.time2
FROM
       mytime
where
       mytime.time1 = mytime.time2
       and mytime.id > 2;

SELECT
       mytime.time1,
       mytime.time2
FROM
       mytime
where
       mytime.time1 > mytime.time2
       and mytime.id = 2;

SELECT
       mytime.time1,
       mytime.time2
FROM
       mytime
where
       mytime.time1 > mytime.time2;

SELECT
       mytime.id,
       mytime.time1,
       mytime.time2
FROM
       mytime
where
       mytime.time1 > mytime.time2;

SELECT
       mytime.id,
       mytime.time1,
       mytime.time2
FROM
       mytime
where
       mytime.time1 > '17:05 :00';

SELECT
       mytime.id,
       mytime.time1,
       mytime.time2
FROM
       mytime
where
       mytime.time1 > '17:05 :00'
       AND mytime.id = 3;