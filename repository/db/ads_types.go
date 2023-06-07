package db

import (
	"dynapgen/constants"
	"time"
)

type AppAds struct {
	ID               string           `db:"id"`
	AppID            string           `db:"app_id"`
	Type             constants.AdType `db:"type"`
	OpenAdID         string           `db:"open_ad_id"`
	BannerAdID       string           `db:"banner_ad_id"`
	InterstitialAdID string           `db:"interstitial_ad_id"`
	RewardAdID       string           `db:"reward_ad_id"`
	NativeAdID       string           `db:"native_ad_id"`
	Order            uint8            `db:"order"`
	CreatedAt        time.Time        `db:"created_at"`
	UpdatedAt        *time.Time       `db:"updated_at"`
}
