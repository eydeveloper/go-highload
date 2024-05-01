CREATE TABLE IF NOT EXISTS user_posts (
    id uuid DEFAULT gen_random_uuid(),
    author_id uuid NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES users(id)
);