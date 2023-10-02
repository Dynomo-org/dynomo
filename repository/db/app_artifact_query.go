package db

const (
	queryGetAppArtifactByID = `SELECT * FROM app_artifacts WHERE id = $1`
	queryUpsertAppArtifact  = `
		INSERT INTO app_artifacts (
			id,
			app_id,
			owner_id,
			name,
			metadata_encrypted,
			download_url,
			build_status,
			created_at,
			updated_at
		) VALUES (
			:id,
			:app_id,
			:owner_id,
			:name,
			:metadata_encrypted,
			:download_url,
			:build_status,
			:created_at,
			:updated_at 
		) ON CONFLICT (id) DO UPDATE
		SET
			download_url = EXCLUDED.download_url,
			build_status = EXCLUDED.build_status,
			updated_at = :updated_at
	`
)
