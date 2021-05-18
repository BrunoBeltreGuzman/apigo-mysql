CREATE TABLE sessions (
       id int NOT NULL primary key AUTO_INCREMENT comment 'primary key',
       user int,
       token VARCHAR (6000),
       timeAt time,
       timeEx time,
       dateAt date,
       dateEx date
) default charset utf8 comment '';