alter table totm.users
    add ipaddress varchar(15);

alter table totm.users
    add enabled bool default false;

alter table totm.users
    add role varchar(10) default 'standard';