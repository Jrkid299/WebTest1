-- Filename migrations/000001_create_user_table.up.sql

CREATE TABLE IF NOT EXISTS userTable (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    username text NOT NULL,
    email text NOT NULL,
    version integer NOT NULL DEFAULT 1
);