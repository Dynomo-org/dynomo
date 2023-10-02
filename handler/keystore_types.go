package handler

type GenerateStoreParam struct {
	KeystoreName  string `json:"keystore_name"`
	FullName      string `json:"full_name"`
	Organization  string `json:"organization,omitempty"`
	Country       string `json:"country"`
	Alias         string `json:"alias"`
	KeyPassword   string `json:"key_password"`
	StorePassword string `json:"store_password"`
}
