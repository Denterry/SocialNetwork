CREATE TABLE IF NOT EXISTS post_events (
    postID String,
    userID String,
    event String,
    timestamp DateTime DEFAULT now()
) ENGINE = MergeTree()
ORDER BY timestamp;
