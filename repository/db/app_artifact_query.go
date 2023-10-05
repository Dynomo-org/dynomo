package db

const (
	queryGetAppArtifacts = `
		SELECT 
			COUNT(*) OVER() AS total,
			artifact.id as id,
			artifact.name as artifact_name,
			app.name as app_name,
			app.id as app_id,
			app.icon_url as icon_url,
			artifact.download_url,
			artifact.build_status,
			artifact.created_at
		FROM app_artifacts artifact
		INNER JOIN apps app 
		ON artifact.app_id = app.id
		%s
		ORDER BY artifact.created_at DESC
	`

	queryUpsertAppArtifact = `
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
