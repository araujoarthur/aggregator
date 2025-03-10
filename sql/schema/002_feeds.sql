-- +goose Up
CREATE TABLE feeds(
   id uuid primary key,
   created_at timestamp,
   updated_at timestamp,
   name text not null,
   url text unique,
   user_id uuid,
   constraint fk_user_uuid foreign key (user_id) references users(id) on delete cascade
);

-- +goose Down

DROP TABLE feeds;