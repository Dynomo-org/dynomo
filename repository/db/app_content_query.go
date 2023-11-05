package db

const (
	queryGetAppContentsByAppID = `SELECT * FROM app_contents WHERE app_id = $1 ORDER BY created_at DESC`
	queryGetAppContentByID     = `SELECT * FROM app_contents WHERE id = $1`
	queryInsertAppContent      = `
		INSERT INTO app_contents(id, app_id, title, description, content, thumbnail_url)
		VALUES (
			:id,
			:app_id,
			:title,
			:description,
			:content,
			:thumbnail_url
		)
	`
	queryUpdateAppContent = `
		UPDATE app_contents
		SET 
			title = :title,
			description = :description,
			content = :content,
			thumbnail_url = :thumbnail_url,
			updated_at = :updated_at
		WHERE id = :id
	`
	queryDeleteAppContent = `DELETE FROM app_contents WHERE id = $1`
)
