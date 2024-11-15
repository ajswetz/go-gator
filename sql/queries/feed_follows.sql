-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
INSERT INTO feed_follows
    (id, created_at, updated_at, user_id, feed_id)
VALUES
    (
        $1,
        $2,
        $3,
        (SELECT id
        FROM users
        WHERE users.name = $4),
        (SELECT id
        FROM feeds
        WHERE feeds.url = $5)
)
RETURNING *
)
SELECT inserted_feed_follow.id,
    inserted_feed_follow.created_at,
    inserted_feed_follow.updated_at,
    inserted_feed_follow.user_id,
    inserted_feed_follow.feed_id,
    users.name AS user_name,
    feeds.name AS feed_name
FROM inserted_feed_follow
    INNER JOIN users ON inserted_feed_follow.user_id = users.id
    INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT users.name AS user_name,
    feeds.name AS feed_name
FROM feed_follows
    INNER JOIN users ON feed_follows.user_id = users.id
    INNER JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE users.name = $1;


-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE (SELECT id
    FROM users
    WHERE users.name = $1) = feed_follows.user_id
    AND (SELECT id
    FROM feeds
    WHERE feeds.url = $2) = feed_follows.feed_id;