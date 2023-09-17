package handler

type GenerateStoreParam struct {
	AppID         string `json:"app_id"`
	FullName      string `json:"full_name"`
	Organization  string `json:"organization,omitempty"`
	Country       string `json:"country"`
	Alias         string `json:"alias"`
	KeyPassword   string `json:"key_password"`
	StorePassword string `json:"store_password"`
}