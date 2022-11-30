CREATE USER customer_user WITH PASSWORD 'customer_pass';

CREATE DATABASE customer;
\c customer;
CREATE SCHEMA customer;
REVOKE ALL ON DATABASE customer FROM PUBLIC;
GRANT ALL PRIVILEGES ON DATABASE customer TO customer_user;
GRANT ALL PRIVILEGES ON SCHEMA customer TO customer_user;
