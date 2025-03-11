-- name: CreatePost :one
INSERT INTO posts(
   id,
   title,
   url,
   description,
   published_at,
   feed_id
) VALUES (
   $1,
   $2,
   $3,
   $4,
   $5,
   $6
) RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.* FROM posts
INNER JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1 ORDER BY posts.published_at DESC LIMIT $2;