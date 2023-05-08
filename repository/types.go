package repository

import (
	"time"
)

type NewAppRequest struct {
	AppName     string
	PackageName string
}

type App struct {
	AppID             string        `json:"id,omitempty"`
	AppName           string        `json:"name,omitempty"`
	AppPackageName    string        `json:"package_name,omitempty"`
	VersionCode       uint          `json:"version_code,omitempty"`
	VersionName       string        `json:"version_name,omitempty"`
	IconURL           string        `json:"icon_url,omitempty"`
	PrivacyPolicyLink string        `json:"privacy_policy_link,omitempty"`
	AdmobAppID        string        `json:"admob_app_id,omitempty"`
	AppLovinSDKKey    string        `json:"applovin_sdk_key,omitempty"`
	AdsConfig         AdsConfig     `json:"ads_config,omitempty"`
	AppConfig         AppConfig     `json:"app_config,omitempty"`
	Contents          []AppContent  `json:"contents,omitempty"`
	Categories        []AppCategory `json:"categories,omitempty"`
	CreatedAt         time.Time     `json:"created_at,omitempty"`
	UpdatedAt         *time.Time    `json:"updated_at,omitempty"`
}

type AdsConfig struct {
	EnableOpenAd         bool `json:"enable_open_ad,omitempty"`
	EnableBannerAd       bool `json:"enable_banner_ad,omitempty"`
	EnableInterstitialAd bool `json:"enable_interstitial_ad,omitempty"`
	EnableRewardAd       bool `json:"enable_reward_ad,omitempty"`
	EnableNativeAd       bool `json:"enable_native_ad,omitempty"`

	PrimaryAd   Ad `json:"primary_ad,omitempty"`
	SecondaryAd Ad `json:"secondary_ad,omitempty"`
	TertiaryAd  Ad `json:"tertiary_ad,omitempty"`

	InterstitialIntervalSecond int      `json:"interstitial_interval_second,omitempty"`
	TestDevices                []string `json:"test_devices,omitempty"`
}

type Ad struct {
	AdType           uint8  `json:"ad_type,omitempty"`
	OpenAdID         string `json:"open_ad_id,omitempty"`
	BannerAdID       string `json:"banner_ad_id,omitempty"`
	InterstitialAdID string `json:"interstitial_ad_id,omitempty"`
	RewardAdID       string `json:"reward_ad_id,omitempty"`
	NativeAdID       string `json:"native_ad_id,omitempty"`
}

type AppCategory struct {
	ID       string `json:"id,omitempty"`
	Category string `json:"category,omitempty"`
}

type AppContent struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Content     string `json:"content,omitempty"`
	CategoryID  string `json:"category_id,omitempty"`
}

type AppConfig struct {
	Strings AppString `json:"strings,omitempty"`
	Style   AppStyle  `json:"style,omitempty"`
}

type AppString struct {
	SetAsWallpaper      string `json:"set_as_wallpaper,omitempty"`
	SetWallpaperHome    string `json:"set_wallpaper_home,omitempty"`
	SetWallpaperLock    string `json:"set_wallpaper_lock,omitempty"`
	WallpaperBoth       string `json:"wallpaper_both,omitempty"`
	Cancel              string `json:"cancel,omitempty"`
	SuccessSetWallpaper string `json:"success_set_wallpaper,omitempty"`
	ExitPromptMessage   string `json:"exit_message,omitempty"`
	NoConnectionMessage string `json:"no_connection_message,omitempty"`
	PrivacyPolicyText   string `json:"privacy_policy_text,omitempty"`
}

type AppStyle struct {
	ColorPrimary          string `json:"color_primary,omitempty"`
	ColorPrimaryVariant   string `json:"color_primary_variant,omitempty"`
	ColorOnPrimary        string `json:"color_on_primary,omitempty"`
	ColorSecondary        string `json:"color_secondary,omitempty"`
	ColorSecondaryVariant string `json:"color_secondary_variant,omitempty"`
	ColorOnSecondary      string `json:"color_on_secondary,omitempty"`
}

type UploadFileParam struct {
	FilePathLocal         string
	FileName              string
	DestinationFolderPath string
	ReplaceIfNameExists   bool
}
