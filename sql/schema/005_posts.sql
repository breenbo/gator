-- +goose Up
CREATE TABLE posts (
id TEXT PRIMARY KEY,
title TEXT NOT NULL,
url TEXT NOT NULL UNIQUE,
description TEXT,
created_at TIMESTAMP,
updated_at TIMESTAMP,
published_at TIMESTAMP,
feed_id TEXT,

CONSTRAINT fk_feed_id
FOREIGN KEY(feed_id) REFERENCES feeds(id)
);
