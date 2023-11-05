package db

const (
	queryDeleteApp       = `DELETE FROM apps WHERE app_id = $1`
	queryGetAppsByUserID = `SELECT count(*) over() as total, * FROM apps WHERE owner_id = $1 ORDER BY created_at DESC`
	queryGetApp          = `SELECT * FROM apps WHERE id = $1`
	queryInsertApp       = `
		insert into apps(
			id,
			owner_id,
			name,
			package_name,
			admob_app_id,
			app_lovin_sdk_key,
			version_name,
			version_code,
			icon_url,
			privacy_policy_link,
			app_strings,
			app_styles,
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
			:admob_app_id,
			:app_lovin_sdk_key,
			:version_name,
			:version_code,
			:icon_url,
			:privacy_policy_link,
			:app_strings,
			:app_styles,
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
			admob_app_id = :admob_app_id,
			app_lovin_sdk_key = :app_lovin_sdk_key,
			version_name = :version_name,
			version_code = :version_code,
			icon_url = :icon_url,
			privacy_policy_link = :privacy_policy_link,
			app_strings = :app_strings,
			app_styles = :app_styles,
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
