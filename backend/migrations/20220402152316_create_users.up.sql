CREATE TABLE users (
    id bigserial not null primary key,
    name varchar not null,
    phone varchar not null unique
);