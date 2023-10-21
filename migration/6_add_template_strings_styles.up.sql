ALTER TABLE templates
ADD COLUMN IF NOT EXISTS strings jsonb NOT NULL DEFAULT '{}'::jsonb,
ADD COLUMN IF NOT EXISTS styles jsonb NOT NULL DEFAULT '{}'::jsonb,
ADD COLUMN IF NOT EXISTS created_at timestamp NOT NULL DEFAULT now(),
ADD COLUMN IF NOT EXISTS updated_at timestamp NULL;