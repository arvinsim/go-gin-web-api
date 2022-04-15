#!/bin/bash

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER newuser;
    CREATE DATABASE newdb;
    GRANT ALL PRIVILEGES ON DATABASE newdb TO newuser;
EOSQL

psql -f /db-dumps/db.dump.sql newdb