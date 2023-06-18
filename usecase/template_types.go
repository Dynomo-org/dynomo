package usecase

type Template struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	RepositoryURL string `json:"repository_url"`
	Type          int    `json:"type"`
}
