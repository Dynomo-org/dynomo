package redis

type BuildStatus uint8

const (
	BuildStatusSuccess BuildStatus = iota + 1
	BuildStatusFail
	BuildStatusInProgress
	BuildStatusPending
)

type Keystore struct {
	Status       BuildStatus `json:"status"`
	URL          string      `json:"url"`
	ErrorMessage string      `json:"error_message"`
}
