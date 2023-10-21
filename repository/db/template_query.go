package db

const (
	queryGetAllTemplates = "select * from templates order by created_at desc"
	queryGetTemplateByID = "select * from templates where id = $1"
	queryInsertTemplate  = `
		insert into templates (
			id,
			name,
			repository_url,
			strings,
			styles,
			type
		) values (
			:id,
			:name,
			:repository_url,
			:strings,
			:styles,
			:type
		)
	`
	queryUpdateTemplate = `
		update templates
		set
			name = :name,
			repository_url = :repository_url,
			strings = :strings,
			styles = :styles,
			type = :type,
			updated_at = :updated_at
		where id = :id
	`
)
