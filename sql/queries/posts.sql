-- name: CreatePost :exec
INSERT INTO posts (id, title, url, description, created_at, updated_at, published_at, feed_id)
VALUES ($1, $2, $3, $4,  $5, $6, $7, $8);

-- name: GetPostForUser :many
SELECT * from posts
WHERE feed_id IN (
    SELECT feed_id FROM feed_follows WHERE user_id = (
        SELECT id FROM users WHERE name = $1
    )
)
ORDER BY published_at DESC
LIMIT $2;
