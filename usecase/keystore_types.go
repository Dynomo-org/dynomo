package usecase

import (
	"dynapgen/repository/db"
	"time"
)

type BuildKeystoreParam struct {
	OwnerID       string
	KeystoreName  string
	FullName      string
	Organization  string
	Country       string
	Alias         string
	KeyPassword   string
	StorePassword string
}

type GetKeystoreListParam struct {
	Page    int
	PerPage int
	OwnerID string
}

type GetKeystoreListResponse struct {
	TotalPage int        `json:"total_page"`
	Page      int        `json:"page"`
	Keystores []Keystore `json:"keystores"`
}

type Keystore struct {
	Total       int       `json:"-"`
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Alias       string    `json:"alias,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	DownloadURL string    `json:"download_url,omitempty"`
	BuildStatus int       `json:"build_status,omitempty"`
}

type keystoreMetadata struct {
	Alias         string `json:"alias,omitempty"`
	FullName      string `json:"full_name,omitempty"`
	Organization  string `json:"organization,omitempty"`
	Country       string `json:"country,omitempty"`
	KeyPassword   string `json:"key_password,omitempty"`
	StorePassword string `json:"store_password,omitempty"`
}

func convertKeystoreFromDB(keystore db.Keystore) Keystore {
	return Keystore{
		Total:       keystore.Total,
		ID:          keystore.ID,
		Name:        keystore.Name,
		Alias:       keystore.Alias,
		CreatedAt:   keystore.CreatedAt,
		DownloadURL: keystore.DownloadURL,
		BuildStatus: keystore.BuildStatus,
	}
}
