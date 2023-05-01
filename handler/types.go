package handler

import "time"

type NewMasterAppRequest struct {
	AppName     string `json:"app_name"`
	PackageName string `json:"package_name"`
}

type MasterApp struct {
	AppID          string        `json:"id,omitempty"`
	AppName        string        `json:"name,omitempty"`
	AppPackageName string        `json:"package_name,omitempty"`
	AdmobAppID     string        `json:"admob_app_id,omitempty"`
	AppLovinSDKKey string        `json:"applovin_sdk_key,omitempty"`
	AdsConfig      AdsConfig     `json:"ads_config,omitempty"`
	AppConfig      AppConfig     `json:"app_config,omitempty"`
	Contents       []AppContent  `json:"contents,omitempty"`
	Categories     []AppCategory `json:"categories,omitempty"`
	CreatedAt      time.Time     `json:"created_at,omitempty"`
	UpdatedAt      *time.Time    `json:"updated_at,omitempty"`
}

type AdsConfig struct {
	EnableOpenAd               bool     `json:"enable_open_ad,omitempty"`
	EnableBannerAd             bool     `json:"enable_banner_ad,omitempty"`
	EnableInterstitialAd       bool     `json:"enable_interstitial_ad,omitempty"`
	EnableRewardAd             bool     `json:"enable_reward_ad,omitempty"`
	EnableNativeAd             bool     `json:"enable_native_ad,omitempty"`
	PrimaryAdType              uint8    `json:"primary_ad_type,omitempty"`
	SecondaryAdType            uint8    `json:"secondary_ad_type,omitempty"`
	TertiaryAdType             uint8    `json:"tiary_ad_type,omitempty"`
	InterstitialIntervalSecond int      `json:"interstitial_interval_second,omitempty"`
	TestDevices                []string `json:"test_devices,omitempty"`
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
	SetWallpaperHome    string `json:"set_wallpaper_home,omitempty"`
	SetWallpaperLock    string `json:"set_wallpaper_lock,omitempty"`
	Cancel              string `json:"cancel,omitempty"`
	SuccessSetWallpaper string `json:"success_set_wallpaper,omitempty"`
	ExitPromptMessage   string `json:"exit_message,omitempty"`
	NoConnectionMessage string `json:"no_connection_message,omitempty"`
	PrivacyPolicyText   string `json:"privacy_policy_text,omitempty"`
}

type AppStyle struct {
	ColorPrimary   string `json:"color_primary,omitempty"`
	ColorSecondary string `json:"color_secondary,omitempty"`
	ColorAccent    string `json:"color_accent,omitempty"`
}
