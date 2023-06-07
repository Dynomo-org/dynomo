package db

const (
	queryDeleteApp             = `delete from apps where app_id = $1`
	queryGetAppsByUserID       = `select count(*) over() as total, * from apps where owner_id = $1`
	queryGetApp                = `select * from apps where id = $1`
	queryGetAppString          = `select * from app_strings where app_id = $1`
	queryGetAppStyle           = `select * from app_styles where app_id = $1`
	queryGetAppContentsByAppID = `select id, title, description, content from app_contents where app_id = $1`
	queryInsertApp             = `
		insert into apps(
			id,
			owner_id,
			name,
			package_name,
			type,
			admob_app_id,
			app_lovin_sdk_key,
			version,
			version_code,
			icon_url,
			privacy_policy_link,
			strings,
			color_primary,
			color_primary_variant,
			color_on_primary,
			enable_open,
			enable_banner,
			enable_interstitial,
			enable_native,
			enable_reward,
			interstitial_interval_second,
			created_at,
			updated_at
		) values (
			:id,
			:owner_id,
			:name,
			:package_name,
			:type,
			:admob_app_id,
			:app_lovin_sdk_key,
			:version,
			:version_code,
			:icon_url,
			:privacy_policy_link,
			:strings,
			:color_primary,
			:color_primary_variant,
			:color_on_primary,
			:enable_open,
			:enable_banner,
			:enable_interstitial,
			:enable_native,
			:enable_reward,
			:interstitial_interval_second,
			:created_at,
			:updated_at
		)
	`
	queryInsertAppAds = `
		insert into app_ads(
			id,
			app_id,
			type,
			open_ad_id,
			banner_ad_id,
			interstitial_ad_id,
			reward_ad_id,
			native_ad_id,
			"order",
			created_at,
			updated_at
		) values (
			:id,
			:app_id,
			:type,
			:open_ad_id,
			:banner_ad_id,
			:interstitial_ad_id,
			:reward_ad_id,
			:native_ad_id,
			:order,
			:created_at,
			:updated_at
		)
	`
	queryUpdateApp = `
		update apps set
			owner_id = :owner_id,
			name = :name,
			package_name = :package_name,
			type = :type,
			admob_app_id = :admob_app_id,
			app_lovin_sdk_key = :app_lovin_sdk_key,
			version = :version,
			version_code = :version_code,
			icon_url = :icon_url,
			privacy_policy_link = :privacy_policy_link,
			strings = :strings,
			color_primary = :color_primary,
			color_primary_variant = :color_primary_variant,
			color_on_primary = :color_on_primary,
			enable_open = :enable_open,
			enable_banner = :enable_banner,
			enable_interstitial = :enable_interstitial,
			enable_native = :enable_native,
			enable_reward = :enable_reward,
			interstitial_interval_second = :interstitial_interval_second,
			created_at = :created_at,
			updated_at = :updated_at
		 where id = :id
	`
)
