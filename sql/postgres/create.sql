CREATE TABLE DRIFT_MIGRATIONS (
    id serial not null primary key,
    name varchar(128) unique not null,
    checksum bytea unique not null,
    applied timestamptz default now()
);