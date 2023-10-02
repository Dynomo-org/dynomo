package nsq

type AppStyle struct {
	ColorPrimary     string `json:"color_primary,omitempty"`
	ColorSecondary   string `json:"color_secondary,omitempty"`
	ColorOnPrimary   string `json:"color_on_primary,omitempty"`
	ColorOnSecondary string `json:"color_on_secondary,omitempty"`
}

type BuildAppParam struct {
	BuildID        string   `json:"build_id,omitempty"`
	AppName        string   `json:"app_name,omitempty"`
	AppVersionCode int      `json:"app_version_code,omitempty"`
	AppVersionName string   `json:"app_version_name,omitempty"`
	AppString      string   `json:"app_string,omitempty"`
	IconUrl        string   `json:"icon_url,omitempty"`
	KeystoreUrl    string   `json:"keystore_url,omitempty"`
	Style          AppStyle `json:"style,omitempty"`
	TemplateType   int      `json:"template_type,omitempty"` // will determine what AppString will be parsed as
	TemplateName   string   `json:"template_name,omitempty"` // example value: simple-wallpaper-base
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
