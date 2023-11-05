package db

import (
	"database/sql"
	"time"
)

type App struct {
	Total                      int          `db:"total"`
	ID                         string       `db:"id"`
	OwnerID                    string       `db:"owner_id"`
	Name                       string       `db:"name"`
	PackageName                string       `db:"package_name"`
	TemplateID                 string       `db:"template_id"`
	AdmobAppID                 string       `db:"admob_app_id"`
	AppLovinSDKKey             string       `db:"app_lovin_sdk_key"`
	Version                    string       `db:"version_name"`
	VersionCode                int          `db:"version_code"`
	IconURL                    string       `db:"icon_url"`
	PrivacyPolicyLink          string       `db:"privacy_policy_link"`
	Strings                    string       `db:"app_strings"` // jsonb type
	Styles                     string       `db:"app_styles"`  // jsonb type
	EnableOpen                 bool         `db:"enable_open"`
	EnableBanner               bool         `db:"enable_banner"`
	EnableInterstitial         bool         `db:"enable_interstitial"`
	EnableNative               bool         `db:"enable_native"`
	EnableReward               bool         `db:"enable_reward"`
	InterstitialIntervalSecond int          `db:"interstitial_interval_second"`
	CreatedAt                  time.Time    `db:"created_at"`
	UpdatedAt                  sql.NullTime `db:"updated_at"`
}
