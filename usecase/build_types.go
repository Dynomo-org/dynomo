package usecase

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

type UpdateBuildStatusParam struct {
	BuildID string
	BuildStatus
}

type BuildStatus struct {
	Status       BuildStatusEnum `json:"status"`
	URL          string          `json:"url"`
	ErrorMessage string          `json:"error_message"`
}
