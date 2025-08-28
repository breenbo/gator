-- +goose Up
CREATE TABLE feed_follows (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id TEXT NOT NULL,
    feed_id TEXT NOT NULL,

    CONSTRAINT fk_user_id
    FOREIGN KEY(user_id) REFERENCES users(id)
    ON DELETE CASCADE,
    
    CONSTRAINT fk_feed_id
    FOREIGN KEY(feed_id) REFERENCES feeds(id)
);

-- +goose Down
DROP TABLE feed_follows;
