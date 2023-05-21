package db

import "time"

type App struct {
	Total                      int         `db:"total"`
	ID                         string      `db:"id"`
	OwnerID                    string      `db:"owner_id"`
	Name                       string      `db:"name"`
	PackageName                string      `db:"package_name"`
	Type                       uint8       `db:"type"`
	AdmobAppID                 string      `db:"admob_app_id"`
	AppLovinSDKKey             string      `db:"app_lovin_sdk_key"`
	Version                    int         `db:"version"`
	VersionCode                string      `db:"version_code"`
	IconURL                    string      `db:"icon_url"`
	PrivacyPolicyLink          string      `db:"privacy_policy_link"`
	Strings                    interface{} `db:"strings"`
	ColorPrimary               string      `db:"color_primary"`
	ColorPrimaryVariant        string      `db:"color_primary_variant"`
	ColorOnPrimary             string      `db:"color_on_primary"`
	EnableOpen                 bool        `db:"enable_open"`
	EnableBanner               bool        `db:"enable_banner"`
	EnableInterstitial         bool        `db:"enable_interstitial"`
	EnableNative               bool        `db:"enable_native"`
	EnableReward               bool        `db:"enable_reward"`
	InterstitialIntervalSecond int         `db:"interstitial_interval_second"`
	CreatedAt                  time.Time   `db:"created_at"`
	UpdatedAt                  *time.Time  `db:"updated_at"`
}
