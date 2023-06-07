package redis

import (
	"dynapgen/constants"
	"time"
)

type AppFull struct {
	ID                string       `json:"id,omitempty"`
	OwnerID           string       `json:"owner_id,omitempty"`
	Name              string       `json:"name,omitempty"`
	PackageName       string       `json:"package_name,omitempty"`
	Type              uint8        `json:"type,omitempty"`
	AdmobAppID        string       `json:"admob_app_id,omitempty"`
	AppLovinSDKKey    string       `json:"app_lovin_sdk_key,omitempty"`
	Version           int          `json:"version,omitempty"`
	VersionCode       string       `json:"version_code,omitempty"`
	IconURL           string       `json:"icon_url,omitempty"`
	PrivacyPolicyLink string       `json:"privacy_policy_link,omitempty"`
	AppConfig         AppConfig    `json:"app_config,omitempty"`
	AdsConfig         AdsConfig    `json:"ads_config,omitempty"`
	Contents          []AppContent `json:"contents,omitempty"`
	CreatedAt         time.Time    `json:"created_at,omitempty"`
	UpdatedAt         *time.Time   `json:"updated_at,omitempty"`
}

type AdsConfig struct {
	EnableOpen         bool `json:"enable_open,omitempty"`
	EnableBanner       bool `json:"enable_banner,omitempty"`
	EnableInterstitial bool `json:"enable_interstitial,omitempty"`
	EnableReward       bool `json:"enable_reward,omitempty"`
	EnableNative       bool `json:"enable_native,omitempty"`

	Ads []AppAd `json:"ads,omitempty"`

	InterstitialIntervalSecond int `json:"interstitial_interval_second,omitempty"`
}

type AppAd struct {
	Type             constants.AdType `json:"ad_type,omitempty"`
	OpenAdID         string           `json:"open_ad_id,omitempty"`
	BannerAdID       string           `json:"banner_ad_id,omitempty"`
	InterstitialAdID string           `json:"interstitial_ad_id,omitempty"`
	RewardAdID       string           `json:"reward_ad_id,omitempty"`
	NativeAdID       string           `json:"native_ad_id,omitempty"`
}

type AppContent struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Content     string `json:"content,omitempty"`
}

type AppConfig struct {
	Strings interface{} `json:"strings,omitempty"`
	Style   AppStyle    `json:"style,omitempty"`
}

type AppStyle struct {
	ColorPrimary        string `json:"color_primary,omitempty"`
	ColorPrimaryVariant string `json:"color_primary_variant,omitempty"`
	ColorOnPrimary      string `json:"color_on_primary,omitempty"`
}
