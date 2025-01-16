create table person (
    id int not null,
    name varchar(100) not null
);

create table users (
    id varchar(36) not null,
    username varchar(100) not null,
    password varchar(100) not null,
    salt varchar(36) not null
);
