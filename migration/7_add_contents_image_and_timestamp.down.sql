ALTER TABLE app_contents
DROP COLUMN IF EXISTS thumbnail_url,
DROP COLUMN IF EXISTS created_at,
DROP COLUMN IF EXISTS updated_at;