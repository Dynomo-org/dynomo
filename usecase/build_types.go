package usecase

type BuildAppParam struct {
	AppID        string
	VersionCode  int
	VersionName  string
	KeystorePath string
}

type BuildStatusEnum int

const (
	BuildStatusEnumSuccess BuildStatusEnum = iota + 1
	BuildStatusEnumFailed
	BuildStatusEnumInProgress
)

type UpdateBuildStatusParam struct {
	AppID string
	BuildStatus
}

type BuildStatus struct {
	Status       BuildStatusEnum `json:"status"`
	URL          string          `json:"url"`
	ErrorMessage string          `json:"error_message"`
}