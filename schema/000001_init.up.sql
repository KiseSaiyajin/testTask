CREATE table users(
    id serial not null unique ,
    name varchar(256) not null,
    username varchar(256) not null unique ,
    password_hash varchar(256) not null
)