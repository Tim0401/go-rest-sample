#!/bin/bash
psql -U user1 -d go << "EOSQL"
drop table IF EXISTS users;
create table users(
  id  SERIAL PRIMARY KEY,
  username varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  created_at timestamp with time zone default CURRENT_TIMESTAMP,
  updated_at timestamp with time zone default CURRENT_TIMESTAMP
);
EOSQL