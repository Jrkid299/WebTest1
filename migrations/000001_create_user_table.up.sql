-- Filename migrations/000001_create_user_table.up.sql

CREATE TABLE IF NOT EXISTS userTable (
    id bigserial PRIMARY KEY,
    username text NOT NULL,
    email text NOT NULL
);