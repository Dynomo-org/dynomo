package handler

type BuildStatusPayload struct {
	Status       int    `json:"status"`
	URL          string `json:"url,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}
