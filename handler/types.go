package handler

var (
	defaultSuccessResponse = map[string]interface{}{
		"success": true,
	}
)

type NewAppRequest struct {
	AppName string `json:"app_name"`
}
