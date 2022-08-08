
/*
  to run script use:
  psql -U postgres -d postgres -a -f .\create_db.sql
*/

CREATE DATABASE darolfonoon
      WITH
      OWNER = postgres
      TEMPLATE = template0
      ENCODING = 'UTF8'
      LC_COLLATE = 'fa_IR.UTF8'
      LC_CTYPE = 'fa_IR.UTF8'
      TABLESPACE = pg_default
      CONNECTION LIMIT = -1;

ALTER DATABASE darolfonoon SET default_transaction_isolation TO 'read committed';

CREATE ROLE darolfonoon_services WITH
    LOGIN
    NOSUPERUSER
    INHERIT
    NOCREATEDB
    NOCREATEROLE
    NOREPLICATION
    PASSWORD 'services';

ALTER Role darolfonoon_services SET default_transaction_isolation TO 'read committed';