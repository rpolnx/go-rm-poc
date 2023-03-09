CREATE USER crud_sql WITH PASSWORD 'crud-password';

CREATE DATABASE 'crud-sql';
\c customer;
CREATE SCHEMA 'crud-sql';
REVOKE ALL ON DATABASE customer FROM PUBLIC;
GRANT ALL PRIVILEGES ON DATABASE customer TO customer_user;
GRANT ALL PRIVILEGES ON SCHEMA customer TO customer_user;
