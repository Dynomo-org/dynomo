package repository

type MasterApp struct {
	Name       string        `json:"name,omitempty"`
	AppConfig  AppConfig     `json:"app_config,omitempty"`
	Contents   []AppContent  `json:"contents,omitempty"`
	Categories []AppCategory `json:"categories,omitempty"`
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
