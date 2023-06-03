package usecase

import (
	"dynapgen/constants"
	"dynapgen/repository/db"
	"dynapgen/repository/redis"
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

type NewAppAdsRequest struct {
	AppID            string
	Type             constants.AdType
	OpenAdID         string
	BannerAdID       string
	InterstitialAdID string
	RewardAdID       string
	NativeAdID       string
	Order            uint8
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

type AppAds struct {
	ID               string           `json:"id"`
	AppID            string           `json:"-"`
	Type             constants.AdType `json:"type"`
	OpenAdID         string           `json:"open_ad_id"`
	BannerAdID       string           `json:"banner_ad_id"`
	InterstitialAdID string           `json:"interstitial_ad_id"`
	RewardAdID       string           `json:"reward_ad_id"`
	NativeAdID       string           `json:"native_ad_id"`
	Order            uint8            `json:"order"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        *time.Time       `json:"updated_at"`
}

type AppFull struct {
	ID                string       `json:"id"`
	OwnerID           string       `json:"owner_id"`
	Name              string       `json:"name"`
	PackageName       string       `json:"package_name"`
	Type              uint8        `json:"type"`
	AdmobAppID        string       `json:"admob_app_id"`
	AppLovinSDKKey    string       `json:"app_lovin_sdk_key"`
	Version           int          `json:"version"`
	VersionCode       string       `json:"version_code"`
	IconURL           string       `json:"icon_url"`
	PrivacyPolicyLink string       `json:"privacy_policy_link"`
	AppConfig         AppConfig    `json:"app_config"`
	AdsConfig         AdsConfig    `json:"ads_config"`
	Content           []AppContent `json:"content"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         *time.Time   `json:"updated_at"`
}

type AdsConfig struct {
	EnableOpen                 bool    `json:"enable_open"`
	EnableBanner               bool    `json:"enable_banner"`
	EnableInterstitial         bool    `json:"enable_interstitial"`
	EnableNative               bool    `json:"enable_native"`
	EnableReward               bool    `json:"enable_reward"`
	InterstitialIntervalSecond int     `json:"interstitial_interval_second"`
	Ads                        []AppAd `json:"ads"`
}

type AppConfig struct {
	Strings interface{} `json:"strings"`
	Style   AppStyle    `json:"style"`
}

type AppStyle struct {
	ColorPrimary        string `json:"color_primary"`
	ColorPrimaryVariant string `json:"color_primary_variant"`
	ColorOnPrimary      string `json:"color_on_primary"`
}

type AppAd struct {
	Type             constants.AdType `json:"type"`
	OpenAdID         string           `json:"open_ad_id"`
	BannerAdID       string           `json:"banner_ad_id"`
	InterstitialAdID string           `json:"interstitial_ad_id"`
	RewardAdID       string           `json:"reward_ad_id"`
	NativeAdID       string           `json:"native_ad_id"`
}

type AppContent struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
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

func buildAppFull(app db.App, appAds []db.AppAds, appContents []db.AppContent) AppFull {
	result := AppFull{
		ID:                app.ID,
		OwnerID:           app.OwnerID,
		Name:              app.Name,
		PackageName:       app.PackageName,
		Type:              app.Type,
		AdmobAppID:        app.AdmobAppID,
		AppLovinSDKKey:    app.AppLovinSDKKey,
		Version:           app.Version,
		VersionCode:       app.VersionCode,
		IconURL:           app.IconURL,
		PrivacyPolicyLink: app.PrivacyPolicyLink,
		AppConfig: AppConfig{
			Strings: app.Strings,
			Style: AppStyle{
				ColorPrimary:        app.ColorPrimary,
				ColorPrimaryVariant: app.ColorPrimaryVariant,
				ColorOnPrimary:      app.ColorOnPrimary,
			},
		},
		AdsConfig: AdsConfig{
			EnableOpen:                 app.EnableOpen,
			EnableBanner:               app.EnableBanner,
			EnableInterstitial:         app.EnableInterstitial,
			EnableNative:               app.EnableNative,
			EnableReward:               app.EnableReward,
			InterstitialIntervalSecond: app.InterstitialIntervalSecond,
		},
		CreatedAt: app.CreatedAt,
		UpdatedAt: app.UpdatedAt,
	}

	ads := make([]AppAd, len(appAds))
	for index, ad := range appAds {
		ads[index] = AppAd{
			Type:             ad.Type,
			OpenAdID:         ad.OpenAdID,
			BannerAdID:       ad.BannerAdID,
			InterstitialAdID: ad.InterstitialAdID,
			RewardAdID:       ad.RewardAdID,
			NativeAdID:       ad.NativeAdID,
		}
	}
	result.AdsConfig.Ads = ads

	contents := make([]AppContent, len(appContents))
	for index, content := range appContents {
		contents[index] = AppContent(content)
	}
	result.Content = contents

	return result
}

func convertAppFullFromCache(cached redis.App) AppFull {
	result := AppFull{
		ID:                cached.ID,
		OwnerID:           cached.OwnerID,
		Name:              cached.Name,
		PackageName:       cached.PackageName,
		Type:              cached.Type,
		AdmobAppID:        cached.AdmobAppID,
		AppLovinSDKKey:    cached.AppLovinSDKKey,
		Version:           cached.Version,
		VersionCode:       cached.VersionCode,
		IconURL:           cached.IconURL,
		PrivacyPolicyLink: cached.PrivacyPolicyLink,
		AppConfig: AppConfig{
			Strings: cached.AppConfig.Strings,
			Style:   AppStyle(cached.AppConfig.Style),
		},
	}

	ads := make([]AppAd, len(cached.AdsConfig.Ads))
	for index, ad := range cached.AdsConfig.Ads {
		ads[index] = AppAd(ad)
	}

	result.AdsConfig = AdsConfig{
		EnableOpen:                 cached.AdsConfig.EnableOpenAd,
		EnableBanner:               cached.AdsConfig.EnableBannerAd,
		EnableInterstitial:         cached.AdsConfig.EnableInterstitialAd,
		EnableNative:               cached.AdsConfig.EnableNativeAd,
		EnableReward:               cached.AdsConfig.EnableRewardAd,
		InterstitialIntervalSecond: cached.AdsConfig.InterstitialIntervalSecond,
		Ads:                        ads,
	}

	return result
}
