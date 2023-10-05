package nsq

type BuildAppParam struct {
	BuildID        string            `json:"build_id,omitempty"`
	AppID          string            `json:"app_id,omitempty"`
	AppName        string            `json:"app_name,omitempty"`
	AppVersionCode int               `json:"app_version_code,omitempty"`
	AppVersionName string            `json:"app_version_name,omitempty"`
	AppStrings     map[string]string `json:"app_strings,omitempty"`
	AppStyles      map[string]string `json:"app_styles,omitempty"`
	IconUrl        string            `json:"icon_url,omitempty"`
	KeystoreUrl    string            `json:"keystore_url,omitempty"`
	TemplateName   string            `json:"template_name,omitempty"` // example value: simple-wallpaper-base
}

type BuildKeystoreParam struct {
	BuildID       string `json:"build_id,omitempty"`
	FullName      string `json:"full_name,omitempty"`
	Organization  string `json:"organization,omitempty"`
	Country       string `json:"country,omitempty"`
	Alias         string `json:"alias,omitempty"`
	KeyPassword   string `json:"key_password,omitempty"`
	StorePassword string `json:"store_password,omitempty"`
}
