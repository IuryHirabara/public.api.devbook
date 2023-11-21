CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS publications;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    createdAt timestamp default current_timestamp() not null
) ENGINE = INNODB;

CREATE TABLE followers(
    userId int not null,
    FOREIGN KEY (userId)
    REFERENCES users(id)
    ON DELETE CASCADE,

    followerId int not null,
    FOREIGN KEY (followerId)
    REFERENCES users(id)
    ON DELETE CASCADE,

    primary key(userId, followerId)
) ENGINE=INNODB;

CREATE TABLE publications(
    id int auto_increment primary key,
    title varchar(50) not null,
    content varchar(300) not null,
    
    authorId int not null,
    FOREIGN KEY (authorId)
    REFERENCES users(id)
    ON DELETE CASCADE,

    likes int default 0,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;