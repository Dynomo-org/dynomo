package usecase

import (
	"dynapgen/repository/db"
	"time"
)

type BuildAppParam struct {
	AppID      string
	KeystoreID string
}

type BuildStatusEnum int

const (
	BuildStatusEnumSuccess BuildStatusEnum = iota + 1
	BuildStatusEnumFailed
	BuildStatusEnumInProgress
	BuildStatusEnumPending
)

type BuildArtifactInfo struct {
	Total       int       `json:"-"`
	ID          string    `json:"id"`
	Name        string    `json:"artifact_name"`
	AppName     string    `json:"app_name"`
	AppID       string    `json:"app_id"`
	IconURL     string    `json:"icon_url"`
	DownloadURL string    `json:"download_url"`
	BuildStatus int       `json:"build_status"`
	CreatedAt   time.Time `json:"created_at"`
}

type BuildStatus struct {
	Status       BuildStatusEnum `json:"status"`
	URL          string          `json:"url"`
	ErrorMessage string          `json:"error_message"`
}

type BuildStatusResponse struct {
	Status BuildStatusEnum `json:"status"`
}

type GetBuildArtifactsParam struct {
	Page    int
	PerPage int
	AppID   string
	OwnerID string
}

type GetBuildArtifactsResponse struct {
	TotalPage int                 `json:"total_page"`
	Page      int                 `json:"page"`
	Artifacts []BuildArtifactInfo `json:"artifacts"`
}

type UpdateBuildStatusParam struct {
	BuildID string
	BuildStatus
}

func convertBuildArtifactInfoFromDB(artifactInfo db.AppArtifactInfo) BuildArtifactInfo {
	return BuildArtifactInfo{
		Total:       artifactInfo.Total,
		ID:          artifactInfo.ID,
		Name:        artifactInfo.Name,
		AppName:     artifactInfo.AppName,
		AppID:       artifactInfo.AppID,
		IconURL:     artifactInfo.IconURL,
		DownloadURL: artifactInfo.DownloadURL,
		BuildStatus: artifactInfo.BuildStatus,
		CreatedAt:   artifactInfo.CreatedAt,
	}
}
