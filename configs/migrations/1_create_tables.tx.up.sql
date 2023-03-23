CREATE TABLE crud_sql."users"
(
    id         serial PRIMARY KEY,
    name       varchar(256),
    email      varchar(256) NOT NULL UNIQUE,
    rg         varchar(20)  NOT NULL UNIQUE,
    cpf        varchar(11)  NOT NULL UNIQUE,

    created_at timestamp DEFAULT now(),
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE crud_sql."companies"
(
    id         serial PRIMARY KEY,
    name       varchar(256) NOT NULL,
    cnpj       varchar(14)  NOT NULL UNIQUE,

    created_at timestamp DEFAULT now(),
    updated_at timestamp,
    deleted_at timestamp
);


CREATE TABLE crud_sql."jobs"
(
    id            serial PRIMARY KEY,
    company2_id    int REFERENCES crud_sql.companies (id),
    user2_id       int REFERENCES crud_sql.users (id),
    month_salary  decimal NOT NULL,
    hours_per_day int     NOT NULL,

    created_at    timestamp DEFAULT now(),
    updated_at    timestamp,
    deleted_at    timestamp
);
