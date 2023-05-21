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

type App struct {
	Total                      int         `json:"total"`
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
