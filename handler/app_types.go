package handler

import "time"

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
	Total                      int         `json:"total"`
	ID                         string      `json:"id"`
	OwnerID                    string      `json:"owner_id"`
	Name                       string      `json:"name"`
	PackageName                string      `json:"package_name"`
	TemplateID                 string      `json:"template_id"`
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
