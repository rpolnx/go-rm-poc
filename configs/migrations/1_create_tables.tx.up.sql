CREATE TABLE crud_sql."users"
(
    id         SERIAL primary key,
    name       varchar(256),
    email      varchar(256),

    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);
