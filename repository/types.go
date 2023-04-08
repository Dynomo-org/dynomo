package repository

type MasterApp struct {
	AppID      string        `json:"id,omitempty"`
	Name       string        `json:"name,omitempty"`
	AdsConfig  AdsConfig     `json:"ads_config,omitempty"`
	AppConfig  AppConfig     `json:"app_config,omitempty"`
	Contents   []AppContent  `json:"contents,omitempty"`
	Categories []AppCategory `json:"categories,omitempty"`
}

type AdsConfig struct {
	EnableBanner       bool `json:"enable_banner,omitempty"`
	EnableInterstitial bool `json:"enable_interstitial,omitempty"`
	EnableRewards      bool `json:"enable_rewards,omitempty"`
}

type AppCategory struct {
	ID       string `json:"id,omitempty"`
	Category string `json:"category,omitempty"`
}

type AppContent struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Content     string `json:"content,omitempty"`
	CategoryID  string `json:"category_id,omitempty"`
}

type AppConfig struct {
	AppName           string   `json:"app_name,omitempty"`
	ExitPromptMessage string   `json:"exit_message,omitempty"`
	Style             AppStyle `json:"app_style,omitempty"`
}

type AppStyle struct {
	ColorPrimary   string `json:"color_primary,omitempty"`
	ColorSecondary string `json:"color_secondary,omitempty"`
	ColorAccent    string `json:"color_accent,omitempty"`
}
