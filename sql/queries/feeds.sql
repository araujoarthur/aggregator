-- name: CreateFeed :one
INSERT INTO feeds (
   id,
   created_at,
   updated_at,
   name,
   url,
   user_id
)values(
   $1,
   $2,
   $3,
   $4,
   $5,
   $6
) RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsWithUserInfo :many

SELECT sqlc.embed(feeds), sqlc.embed(users) FROM feeds INNER JOIN users ON users.id = feeds.user_id;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url=$1;

-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at=$1, updated_at=$1 WHERE id=$2;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds ORDER BY last_fetched_at NULLS FIRST LIMIT 1;