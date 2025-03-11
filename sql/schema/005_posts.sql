-- +goose Up
CREATE TABLE posts(
   id uuid primary key,
   created_at timestamp not null default current_timestamp,
   updated_at timestamp not null default current_timestamp,
   title text not null,
   url text not null unique,
   description text,
   published_at timestamp,
   feed_id uuid not null,
   constraint fk_post_feed_id foreign key (feed_id) references feeds(id)
);

-- +goose Down
DROP TABLE posts;