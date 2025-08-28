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
