ALTER TABLE app_contents
ADD COLUMN IF NOT EXISTS thumbnail_url TEXT,
ADD COLUMN IF NOT EXISTS created_at timestamp NOT NULL DEFAULT NOW(),
ADD COLUMN IF NOT EXISTS updated_at timestamp;