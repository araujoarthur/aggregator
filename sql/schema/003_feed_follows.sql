-- +goose Up
CREATE TABLE feed_follows(
   id uuid primary key,
   created_at timestamp,
   updated_at timestamp,
   user_id uuid not null,
   feed_id uuid not null,
   unique(user_id, feed_id),
   constraint fk_user_id foreign key (user_id) references users(id) on delete cascade,
   constraint fk_feed_id foreign key (feed_id) references feeds(id) on delete cascade
);

-- +goose Down
DROP TABLE feed_follows;