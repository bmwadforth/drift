CREATE TABLE DRIFT_MIGRATIONS (
    id serial not null primary key,
    name varchar(128) unique not null,
    checksum varchar(20) unique not null,
    applied timestamptz default now()
);