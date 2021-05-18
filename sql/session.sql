Drop TABLE if EXISTS sessions;

CREATE TABLE `sessions` (
       `id` int NOT NULL AUTO_INCREMENT COMMENT 'primary key',
       `user` int not NULL,
       `token` varchar(6000) DEFAULT NULL,
       `ip` varchar(30) DEFAULT NULL,
       `createdAt` datetime DEFAULT now(),
       `expiredAt` datetime DEFAULT now(),
       PRIMARY KEY (`id`),
       created_timeAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment 'created tiem',
       updated_timeAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment 'updated tiem'
) ENGINE = InnoDB AUTO_INCREMENT = 3 DEFAULT CHARSET = utf8;

alter table
       sessions
add
       foreign key(user) references users(id);