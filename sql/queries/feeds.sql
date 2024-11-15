-- name: CreateFeed :one
INSERT INTO feeds
    (id, created_at, updated_at, name, url, user_id)
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6
)
RETURNING *;

-- name: GetAllFeeds :many
SELECT feeds.name AS feed_name, feeds.url, users.name AS user_name
FROM feeds
    INNER JOIN users
    ON feeds.user_id = users.id;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = $2,
    last_fetched_at = $2
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT
1;