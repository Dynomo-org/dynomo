package usecase

import "time"

type NewMasterAppRequest struct {
	AppName     string
	PackageName string
}

type MasterApp struct {
	AppID          string
	AppName        string
	AppPackageName string
	AdmobAppID     string
	AppLovinSDKKey string
	AdsConfig      AdsConfig
	AppConfig      AppConfig
	Contents       []AppContent
	Categories     []AppCategory
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

type AdsConfig struct {
	EnableOpenAd               bool
	EnableBannerAd             bool
	EnableInterstitialAd       bool
	EnableRewardAd             bool
	EnableNativeAd             bool
	PrimaryAdType              uint8
	SecondaryAdType            uint8
	TertiaryAdType             uint8
	InterstitialIntervalSecond int
	TestDevices                []string
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
	SetWallpaperHome    string
	SetWallpaperLock    string
	Cancel              string
	SuccessSetWallpaper string
	ExitPromptMessage   string
	NoConnectionMessage string
	PrivacyPolicyText   string
}

type AppStyle struct {
	ColorPrimary   string
	ColorSecondary string
	ColorAccent    string
}
