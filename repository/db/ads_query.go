package db

const (
	queryGetAppAdsByAppID = `select * from app_ads where app_id = $1 order by created_at desc`
	queryGetAppAdsByID    = `select * from app_ads where id = $1`
	queryUpdateAppAds     = `
		update app_ads
		SET (
			type,
			open_ad_id,
			banner_ad_id,
			interstitial_ad_id,
			reward_ad_id,
			native_ad_id,
			order,
			updated_at
		) values (
			:type,
			:open_ad_id,
			:banner_ad_id,
			:interstitial_ad_id,
			:reward_ad_id,
			:native_ad_id,
			:order,
			:updated_at
		) where id = :id
	`
)
