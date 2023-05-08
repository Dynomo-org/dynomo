package handler

import "time"

type NewAppRequest struct {
	AppName     string `json:"app_name"`
	PackageName string `json:"package_name"`
}

type App struct {
	AppID             string        `json:"id"`
	AppName           string        `json:"name"`
	AppPackageName    string        `json:"package_name"`
	VersionCode       uint          `json:"version_code"`
	VersionName       string        `json:"version_name"`
	IconURL           string        `json:"icon_url"`
	PrivacyPolicyLink string        `json:"privacy_policy_link"`
	AdmobAppID        string        `json:"admob_app_id"`
	AppLovinSDKKey    string        `json:"applovin_sdk_key"`
	AdsConfig         AdsConfig     `json:"ads_config"`
	AppConfig         AppConfig     `json:"app_config"`
	Contents          []AppContent  `json:"contents"`
	Categories        []AppCategory `json:"categories"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         *time.Time    `json:"updated_at"`
}

type AdsConfig struct {
	EnableOpenAd         bool `json:"enable_open_ad"`
	EnableBannerAd       bool `json:"enable_banner_ad"`
	EnableInterstitialAd bool `json:"enable_interstitial_ad"`
	EnableRewardAd       bool `json:"enable_reward_ad"`
	EnableNativeAd       bool `json:"enable_native_ad"`

	PrimaryAd   Ad `json:"primary_ad"`
	SecondaryAd Ad `json:"secondary_ad"`
	TertiaryAd  Ad `json:"tertiary_ad"`

	InterstitialIntervalSecond int      `json:"interstitial_interval_second"`
	TestDevices                []string `json:"test_devices"`
}

type Ad struct {
	AdType           uint8  `json:"ad_type"`
	OpenAdID         string `json:"open_ad_id"`
	BannerAdID       string `json:"banner_ad_id"`
	InterstitialAdID string `json:"interstitial_ad_id"`
	RewardAdID       string `json:"reward_ad_id"`
	NativeAdID       string `json:"native_ad_id"`
}

type AppCategory struct {
	ID       string `json:"id"`
	Category string `json:"category"`
}

type AppContent struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	CategoryID  string `json:"category_id"`
}

type AppConfig struct {
	Strings AppString `json:"strings"`
	Style   AppStyle  `json:"style"`
}

type AppString struct {
	SetAsWallpaper      string `json:"set_as_wallpaper"`
	SetWallpaperHome    string `json:"set_wallpaper_home"`
	SetWallpaperLock    string `json:"set_wallpaper_lock"`
	WallpaperBoth       string `json:"wallpaper_both"`
	Cancel              string `json:"cancel"`
	SuccessSetWallpaper string `json:"success_set_wallpaper"`
	ExitPromptMessage   string `json:"exit_message"`
	NoConnectionMessage string `json:"no_connection_message"`
	PrivacyPolicyText   string `json:"privacy_policy_text"`
}

type AppStyle struct {
	ColorPrimary          string `json:"color_primary"`
	ColorPrimaryVariant   string `json:"color_primary_variant"`
	ColorOnPrimary        string `json:"color_on_primary"`
	ColorSecondary        string `json:"color_secondary"`
	ColorSecondaryVariant string `json:"color_secondary_variant"`
	ColorOnSecondary      string `json:"color_on_secondary"`
}
