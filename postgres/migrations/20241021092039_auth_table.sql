-- +goose Up
-- 001_create_users_table.up.sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  role TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP -- NOT NULL DEFAULT NOW()
);

-- +goose StatementBegin
-- SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE users
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
