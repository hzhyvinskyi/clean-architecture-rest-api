CREATE TABLE IF NOT EXISTS articles (
    article_id UUID DEFAULT uuid_generate_v4(),
    user_id UUID,
    title VARCHAR NOT NULL,
    description TEXT NOT NULL,
    image_url VARCHAR NOT NULL,
    likes INT DEFAULT 0,
    status BOOLEAN DEFAULT '0',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (article_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
);