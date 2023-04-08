package usecase

type MasterApp struct {
	AppID      string
	Name       string
	AdsConfig  AdsConfig
	AppConfig  AppConfig
	Contents   []AppContent
	Categories []AppCategory
}

type AdsConfig struct {
	EnableBanner       bool
	EnableInterstitial bool
	EnableRewards      bool
}

type AppCategory struct {
	ID       string
	Category string
}

type AppContent struct {
	Title       string
	Description string
	Content     string
	CategoryID  string
}

type AppConfig struct {
	AppName           string
	ExitPromptMessage string
	Style             AppStyle
}

type AppStyle struct {
	ColorPrimary   string
	ColorSecondary string
	ColorAccent    string
}
