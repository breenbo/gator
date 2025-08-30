-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: ListFeed :many
SELECT feeds.name, feeds.url, users.name as userName
FROM feeds
JOIN users ON users.id = feeds.user_id;

-- name: GetFeedIDFromURL :one
SELECT id FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at=$1, last_fetched_at=$1
WHERE id=$2;

-- name: GetNextFeedToFetch :one
SELECT id, url FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
