-- Filename migrations/000002_add_user_check_constraint.up.sql

-- ALTER TABLE user ADD CONSTRAINT mode_length_check CHECK (array_length(mode, 1) BETWEEN 1 AND 5);