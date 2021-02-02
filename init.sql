CREATE TABLE adverts
(
    id serial not null primary key,
    name varchar(200),
    description text,
    photo_links varchar(1000),
    price int,
    creation_date date default now()
);