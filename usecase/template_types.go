package usecase

import (
	"database/sql"
	"dynapgen/repository/db"
	"dynapgen/util/log"
	"encoding/json"
	"time"
)

type Template struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	RepositoryURL string            `json:"repository_url"`
	Styles        map[string]string `json:"styles"`
	Strings       map[string]string `json:"strings"`
	Type          int               `json:"type"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
}

func (t *Template) convertToDB() db.Template {
	if t.Styles == nil {
		t.Styles = make(map[string]string)
	}

	if t.Strings == nil {
		t.Strings = make(map[string]string)
	}

	styles, err := json.Marshal(t.Styles)
	if err != nil {
		log.Error(err, "error marshalling template styles", t.Styles)
	}

	strings, err := json.Marshal(t.Strings)
	if err != nil {
		log.Error(err, "error marshalling template strings", t.Strings)
	}

	result := db.Template{
		ID:            t.ID,
		Name:          t.Name,
		RepositoryURL: t.RepositoryURL,
		Styles:        string(styles),
		Strings:       string(strings),
		Type:          t.Type,
		CreatedAt:     t.CreatedAt,
	}

	if !t.UpdatedAt.IsZero() {
		result.UpdatedAt = sql.NullTime{Valid: true, Time: t.UpdatedAt}
	}

	return result
}

func convertTemplateFromDB(template db.Template) Template {
	result := Template{
		ID:            template.ID,
		Name:          template.Name,
		RepositoryURL: template.RepositoryURL,
		Type:          template.Type,
		CreatedAt:     template.CreatedAt,
	}

	if err := json.Unmarshal([]byte(template.Strings), &result.Strings); err != nil {
		log.Error(err, "error unmarshalling template strings", template.Strings)
	}

	if err := json.Unmarshal([]byte(template.Styles), &result.Styles); err != nil {
		log.Error(err, "error unmarshalling template styles", template.Styles)
	}

	if template.UpdatedAt.Valid {
		result.UpdatedAt = template.UpdatedAt.Time
	}

	return result
}
