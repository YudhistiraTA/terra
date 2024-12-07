CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE EXTENSION IF NOT EXISTS fuzzystrmatch;

CREATE TABLE
  IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    refresh_token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
  );

CREATE TABLE
  IF NOT EXISTS posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW ()
  );

CREATE INDEX IF NOT EXISTS idx_posts_title_trgm ON posts USING gin (title gin_trgm_ops);

CREATE INDEX IF NOT EXISTS idx_posts_content_trgm ON posts USING gin (content gin_trgm_ops);