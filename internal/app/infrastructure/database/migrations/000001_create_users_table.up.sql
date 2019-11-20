CREATE TABLE IF NOT EXISTS users (
    user_id UUID DEFAULT uuid_generate_v4(),
    email VARCHAR NOT NULL,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    avatar VARCHAR DEFAULT NULL,
    is_verified BOOLEAN DEFAULT '0',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id)
);