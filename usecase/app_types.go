package usecase

import (
	"errors"
	"time"
)

const (
	AppTypeUnset     = 0
	AppTypeGuide     = 1
	AppTypeWallpaper = 2
)

var (
	errorAppNotFound = errors.New("app not found")
)

type NewAppRequest struct {
	AppName     string
	PackageName string
	OwnerID     string
}

type App struct {
	Total                      int         `json:"-"`
	ID                         string      `json:"id"`
	OwnerID                    string      `json:"owner_id"`
	Name                       string      `json:"name"`
	PackageName                string      `json:"package_name"`
	Type                       uint8       `json:"type"`
	AdmobAppID                 string      `json:"admob_app_id"`
	AppLovinSDKKey             string      `json:"app_lovin_sdk_key"`
	Version                    int         `json:"version"`
	VersionCode                string      `json:"version_code"`
	IconURL                    string      `json:"icon_url"`
	PrivacyPolicyLink          string      `json:"privacy_policy_link"`
	Strings                    interface{} `json:"strings"`
	ColorPrimary               string      `json:"color_primary"`
	ColorPrimaryVariant        string      `json:"color_primary_variant"`
	ColorOnPrimary             string      `json:"color_on_primary"`
	EnableOpen                 bool        `json:"enable_open"`
	EnableBanner               bool        `json:"enable_banner"`
	EnableInterstitial         bool        `json:"enable_interstitial"`
	EnableNative               bool        `json:"enable_native"`
	EnableReward               bool        `json:"enable_reward"`
	InterstitialIntervalSecond int         `json:"interstitial_interval_second"`
	CreatedAt                  time.Time   `json:"created_at"`
	UpdatedAt                  *time.Time  `json:"updated_at"`
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

type GetAppListParam struct {
	Page    int
	PerPage int
	OwnerID string
}

type GetAppListResponse struct {
	TotalPage int   `json:"total_page"`
	Page      int   `json:"page"`
	Apps      []App `json:"apps"`
}

type WallpaperStrings struct {
	SetAsWallpaper      string `json:"set_as_wallpaper"`
	SetWallpaperHome    string `json:"set_wallpaper_home"`
	SetWallpaperLock    string `json:"set_wallpaper_lock"`
	WallpaperBoth       string `json:"wallpaper_both"`
	Cancel              string `json:"cancel"`
	SuccessSetWallpaper string `json:"success_set_wallpaper"`
	ExitPromptMessage   string `json:"exit_prompt_message"`
	NoConnectionMessage string `json:"no_connection_message"`
	PrivacyPolicyText   string `json:"privacy_policy_text"`
}

// updateWith will update the app with the given input ONLY if the input attribute is not empty
func (app *App) updateWith(input App) {
	if input.Name != "" {
		app.Name = input.Name
	}
	if input.PackageName != "" {
		app.PackageName = input.PackageName
	}
	if input.VersionCode != "" {
		app.VersionCode = input.VersionCode
	}
	if input.Version != 0 {
		app.Version = input.Version
	}
	if input.IconURL != "" {
		app.IconURL = input.IconURL
	}
	if input.PrivacyPolicyLink != "" {
		app.PrivacyPolicyLink = input.PrivacyPolicyLink
	}

	// style settings
	if input.ColorPrimary != "" {
		app.ColorPrimary = input.ColorPrimary
	}
	if input.ColorPrimaryVariant != "" {
		app.ColorPrimaryVariant = input.ColorPrimaryVariant
	}
	if input.ColorOnPrimary != "" {
		app.ColorOnPrimary = input.ColorOnPrimary
	}

	// string settings
	if input.Strings != nil {
		app.Strings = input.Strings
	}

	if input.EnableOpen != app.EnableOpen {
		app.EnableOpen = input.EnableOpen
	}
	if input.EnableBanner != app.EnableBanner {
		app.EnableBanner = input.EnableBanner
	}
	if input.EnableInterstitial != app.EnableInterstitial {
		app.EnableInterstitial = input.EnableInterstitial
	}
	if input.EnableNative != app.EnableNative {
		app.EnableNative = input.EnableNative
	}
	if input.EnableReward != app.EnableReward {
		app.EnableReward = input.EnableReward
	}
	if input.InterstitialIntervalSecond != app.InterstitialIntervalSecond {
		app.InterstitialIntervalSecond = input.InterstitialIntervalSecond
	}
}
