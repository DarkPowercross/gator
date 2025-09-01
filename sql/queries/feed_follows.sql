-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT 
    feed_follow.id,
    feed_follow.created_at,
    feed_follow.updated_at,
    feed_follow.user_id,
    users.name AS user_name,
    feed_follow.feed_id,
    feeds.name AS feed_name
FROM inserted_feed_follow AS feed_follow
JOIN users ON feed_follow.user_id = users.id
JOIN feeds ON feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT 
    feed_follow.id,
    feed_follow.created_at,
    feed_follow.updated_at,
    feed_follow.user_id,
    users.name   AS user_name,
    feed_follow.feed_id,
    feeds.name   AS feed_name,
    feeds.url    AS feed_url
FROM feed_follows AS feed_follow
JOIN users ON feed_follow.user_id = users.id
JOIN feeds ON feed_follow.feed_id = feeds.id
WHERE feed_follow.user_id = $1;

-- name: UnFollow :exec
DELETE FROM feed_follows WHERE user_id = $1 and feed_id = $2;