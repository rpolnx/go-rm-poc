CREATE TABLE crud_sql."users"
(
    id         SERIAL primary key,
    name       varchar(256),
    email      varchar(256),

    created_at timestamp DEFAULT now(),
    updated_at timestamp,
    deleted_at timestamp
);
