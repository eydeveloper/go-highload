CREATE TABLE IF NOT EXISTS user_followers (
    id uuid DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    follower_id uuid NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (follower_id) REFERENCES users(id)
);