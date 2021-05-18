SELECT
       *
FROM
       mydata
where
       date1 = '2021-05-10';

INSERT INTO
       mydata (date1, date2, date3)
VAlUES
       (now(), '2021-05-20', ' 2021-05-25');

SELECT
       mydata.date1,
       mydata.date2
FROM
       mydata
WHERE
       mydata.date1 = mydata.date2
       AND mydata.id = 1;