drop database if exists students;
create database students;
use students;

create table student(
                     id int NOT NULL PRIMARY KEY,
                     name VARCHAR(40) NOT NULL,
                     email VARCHAR(40) NOT NULL UNIQUE,
                     phone VARCHAR(10),
                     age int
);

INSERT INTO student VALUES (1, 'John', 'john21@example.com', '9728810299', 21);
INSERT INTO student VALUES (2, 'Jess', 'jessH90@example.com', '9844537168', 19);
