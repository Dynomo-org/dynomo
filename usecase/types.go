package usecase

import "time"

type NewAppRequest struct {
	AppName     string
	PackageName string
}

type App struct {
	AppID             string
	AppName           string
	AppPackageName    string
	VersionCode       uint
	VersionName       string
	IconURL           string
	PrivacyPolicyLink string
	AdmobAppID        string
	AppLovinSDKKey    string
	AdsConfig         AdsConfig
	AppConfig         AppConfig
	Contents          []AppContent
	Categories        []AppCategory
	CreatedAt         time.Time
	UpdatedAt         *time.Time
}

type AdsConfig struct {
	EnableOpenAd         bool
	EnableBannerAd       bool
	EnableInterstitialAd bool
	EnableRewardAd       bool
	EnableNativeAd       bool

	PrimaryAd   Ad
	SecondaryAd Ad
	TertiaryAd  Ad

	InterstitialIntervalSecond int
	TestDevices                []string
}

type Ad struct {
	AdType           uint8
	OpenAdID         string
	BannerAdID       string
	InterstitialAdID string
	RewardAdID       string
	NativeAdID       string
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
	SetAsWallpaper      string
	SetWallpaperHome    string
	SetWallpaperLock    string
	WallpaperBoth       string
	Cancel              string
	SuccessSetWallpaper string
	ExitPromptMessage   string
	NoConnectionMessage string
	PrivacyPolicyText   string
}

type AppStyle struct {
	ColorPrimary          string
	ColorPrimaryVariant   string
	ColorOnPrimary        string
	ColorSecondary        string
	ColorSecondaryVariant string
	ColorOnSecondary      string
}
