-- +goose Up
CREATE TABLE users(
   id UUID primary key,
   created_at timestamp,
   updated_at timestamp,
   name text not null
);

-- +goose Down
DROP TABLE users;