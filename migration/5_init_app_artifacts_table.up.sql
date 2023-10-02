CREATE TABLE IF NOT EXISTS app_artifacts (
    id VARCHAR,
    app_id VARCHAR,
    owner_id VARCHAR,
    name VARCHAR,
    metadata_encrypted VARCHAR,
    download_url VARCHAR,
    build_status SMALLINT,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
	CONSTRAINT app_artifacts_pkey PRIMARY KEY (id)
);