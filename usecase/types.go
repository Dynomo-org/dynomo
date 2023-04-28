package usecase

type MasterApp struct {
	AppID      string
	AppName    string
	AdsConfig  AdsConfig
	AppConfig  AppConfig
	Contents   []AppContent
	Categories []AppCategory
}

type AdsConfig struct {
	EnableBanner       bool
	EnableInterstitial bool
	EnableReward       bool
}

type AppCategory struct {
	ID       string
	Category string
}

type AppContent struct {
	ID          string
	Title       string
	Description string
	Content     string
	CategoryID  string
}

type AppConfig struct {
	Strings AppString
	Style   AppStyle
}

type AppString struct {
	ExitPromptMessage   string
	NoConnectionMessage string
	PrivacyPolicyText   string
}

type AppStyle struct {
	ColorPrimary   string
	ColorSecondary string
	ColorAccent    string
}
