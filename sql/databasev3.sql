DROP TABLE IF EXISTS roles;

CREATE TABLE roles(
       id int NOT NULL primary key AUTO_INCREMENT comment 'primary key',
       role varchar(255) NOT NULL,
       createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment 'created tiem',
       updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment 'updated tiem'
);

INSERT INTO
       roles (role)
values
       ("admin");

DROP TABLE IF EXISTS users;

CREATE TABLE users(
       id int NOT NULL primary key AUTO_INCREMENT comment 'primary key',
       name varchar(255) NOT NULL,
       email varchar(255) NOT NULL,
       password varchar(255) NOT NULL,
       role int NOT NULL default 1,
       createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment 'created tiem',
       updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment 'updated tiem'
);

DROP TABLE IF EXISTS posts;

CREATE TABLE posts(
       id int NOT NULL primary key AUTO_INCREMENT comment 'primary key',
       user int NOT NULL,
       title varchar(255) NOT NULL,
       content varchar(255) NOT NULL,
       createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment 'created tiem',
       updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP comment 'updated tiem'
);

alter table
       users
add
       foreign key(role) references roles(id);

alter table
       posts
add
       foreign key(user) references users(id);