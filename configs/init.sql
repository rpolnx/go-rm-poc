CREATE USER local_admin WITH PASSWORD 'password';

CREATE DATABASE crud_sql;
\c crud_sql;
CREATE SCHEMA crud_sql;
REVOKE ALL ON DATABASE crud_sql FROM PUBLIC;
GRANT ALL PRIVILEGES ON DATABASE crud_sql TO local_admin;
GRANT ALL PRIVILEGES ON SCHEMA crud_sql TO local_admin;
