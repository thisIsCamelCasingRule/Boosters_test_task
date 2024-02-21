CREATE TABLE IF NOT EXISTS posts (
    id BIGSERIAL PRIMARY KEY,
    title TEXT,
    content TEXT,
    created_at timestamptz,
    updated_at timestamptz
);