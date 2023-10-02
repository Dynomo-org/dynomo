CREATE TABLE IF NOT EXISTS keystores (
    id VARCHAR,
    owner_id VARCHAR,
    name VARCHAR,
    alias VARCHAR,
    metadata_encrypted VARCHAR,
    download_url VARCHAR,
    build_status SMALLINT,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
	CONSTRAINT keystores_pkey PRIMARY KEY (id)
);