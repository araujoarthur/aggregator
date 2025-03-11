-- name: CreateFeedFollow :one
WITH inserted_follow AS (
   INSERT INTO feed_follows(
      id,
      created_at,
      updated_at,
      user_id,
      feed_id
   )
   VALUES(
      $1,
      $2,
      $3,
      $4,
      $5
   ) RETURNING *
) SELECT inserted_follow.*, 
         feeds.name AS feed_name, 
         users.name as user_name 
   FROM inserted_follow 
   INNER JOIN feeds ON feeds.id = inserted_follow.feed_id
   INNER JOIN users on users.id = inserted_follow.user_id;
   

-- name: UserFollowsFeed :one
SELECT EXISTS (
   SELECT 1 FROM feed_follows WHERE user_id = $1 AND feed_id = $2
) as follows;


-- name: GetFollowsByUser :many
SELECT feed_follows.*, feeds.name as feed_name FROM feed_follows 
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1;