package db

const (
	queryGetKeystoreByOwnerID = `SELECT COUNT(*) OVER() AS total, * FROM keystores WHERE owner_id = $1 ORDER BY created_at DESC`
	queryGetKeystoreByID      = `SELECT * FROM keystores WHERE id = $1`
	queryUpsertKeystore       = `
		INSERT INTO keystores(
			id,
			owner_id,
			name,
			alias,
			metadata_encrypted,
			download_url,
			build_status,
			created_at
		) VALUES (
			:id,
			:owner_id,
			:name,
			:alias,
			:metadata_encrypted,
			:download_url,
			:build_status,
			:created_at
		) ON CONFLICT (id) DO UPDATE 
		SET 
			download_url = EXCLUDED.download_url,
			build_status = EXCLUDED.build_status,
			updated_at = :updated_at
	`
)
