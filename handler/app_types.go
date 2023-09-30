package handler

import "dynapgen/usecase"

const (
	defaultPerPage = 10
	defaultPage    = 1
)

type NewAppRequest struct {
	AppName     string `json:"app_name"`
	PackageName string `json:"package_name"`
}

type NewAppAdsRequest struct {
	AppID            string `json:"app_id"`
	Type             uint   `json:"type"`
	OpenAdID         string `json:"open_ad_id"`
	BannerAdID       string `json:"banner_ad_id"`
	InterstitialAdID string `json:"interstitial_ad_id"`
	RewardAdID       string `json:"reward_ad_id"`
	NativeAdID       string `json:"native_ad_id"`
}

type App struct {
	OwnerID                    string            `json:"owner_id"`
	Name                       string            `json:"name"`
	PackageName                string            `json:"package_name"`
	TemplateID                 string            `json:"template_id"`
	AdmobAppID                 string            `json:"admob_app_id"`
	AppLovinSDKKey             string            `json:"app_lovin_sdk_key"`
	Version                    string            `json:"version"`
	VersionCode                int               `json:"version_code"`
	IconURL                    string            `json:"icon_url"`
	PrivacyPolicyLink          string            `json:"privacy_policy_link"`
	Strings                    map[string]string `json:"strings"`
	Styles                     map[string]string `json:"styles"`
	EnableOpen                 bool              `json:"enable_open"`
	EnableBanner               bool              `json:"enable_banner"`
	EnableInterstitial         bool              `json:"enable_interstitial"`
	EnableNative               bool              `json:"enable_native"`
	EnableReward               bool              `json:"enable_reward"`
	InterstitialIntervalSecond int               `json:"interstitial_interval_second"`
}

func (a *App) convertToUsecase() usecase.App {
	return usecase.App{
		OwnerID:                    a.OwnerID,
		Name:                       a.Name,
		PackageName:                a.PackageName,
		TemplateID:                 a.TemplateID,
		AdmobAppID:                 a.AdmobAppID,
		AppLovinSDKKey:             a.AppLovinSDKKey,
		Version:                    a.Version,
		VersionCode:                a.VersionCode,
		IconURL:                    a.IconURL,
		PrivacyPolicyLink:          a.PrivacyPolicyLink,
		Strings:                    a.Strings,
		Styles:                     a.Styles,
		EnableOpen:                 a.EnableOpen,
		EnableBanner:               a.EnableBanner,
		EnableInterstitial:         a.EnableInterstitial,
		EnableNative:               a.EnableNative,
		EnableReward:               a.EnableReward,
		InterstitialIntervalSecond: a.InterstitialIntervalSecond,
	}
}
