create table if not exists totm.blogposts
(
    blogpostid varchar(36) not null,
    title      text,
    markdown   text,
    category   varchar(20),
    image      varchar(50),
    video      varchar(20),
    date       date
);