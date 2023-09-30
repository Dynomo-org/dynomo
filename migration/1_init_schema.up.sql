CREATE TABLE IF NOT EXISTS app_ads (
	id varchar NOT NULL,
	app_id varchar NULL,
	"type" int4 NULL,
	open_ad_id varchar NULL,
	banner_ad_id varchar NULL,
	interstitial_ad_id varchar NULL,
	reward_ad_id varchar NULL,
	native_ad_id varchar NULL,
	"order" int4 NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT app_ads_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS app_contents (
	id varchar NOT NULL,
	app_id varchar NULL,
	title varchar NULL,
	description varchar NULL,
	"content" varchar NULL,
	category_id varchar NULL,
	CONSTRAINT app_contents_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS apps (
	id varchar NOT NULL,
	owner_id varchar NOT NULL DEFAULT '',
	"name" varchar NOT NULL DEFAULT '',
	package_name varchar NOT NULL DEFAULT '',
	version_code int4 NOT NULL DEFAULT 1,
	version_name varchar NOT NULL DEFAULT '1.0',
	icon_url varchar NOT NULL DEFAULT '',
	privacy_policy_link varchar NOT NULL DEFAULT '',
	app_strings jsonb NOT NULL DEFAULT '{}'::jsonb,
	app_styles jsonb NOT NULL DEFAULT '{}'::jsonb,
	admob_app_id varchar NOT NULL DEFAULT '',
	app_lovin_sdk_key varchar NOT NULL DEFAULT '',
	enable_open bool NOT NULL DEFAULT FALSE,
	enable_banner bool NOT NULL DEFAULT FALSE,
	enable_interstitial bool NOT NULL DEFAULT FALSE,
	enable_native bool NOT NULL DEFAULT FALSE,
	enable_reward bool NOT NULL DEFAULT FALSE,
	interstitial_interval_second int4 NOT NULL DEFAULT 60,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	template_id varchar NOT NULL DEFAULT ''::character varying,
	CONSTRAINT apps_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS roles (
	id varchar NOT NULL,
	"name" varchar NULL,
	CONSTRAINT roles_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS templates (
	id varchar NOT NULL,
	"name" varchar NULL,
	repository_url varchar NULL,
	"type" int4 NULL,
	CONSTRAINT templates_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_roles (
	user_id varchar NOT NULL,
	role_id varchar NOT NULL,
	CONSTRAINT user_roles_pkey PRIMARY KEY (user_id, role_id)
);

CREATE TABLE IF NOT EXISTS users (
	id varchar NOT NULL,
	full_name varchar NULL,
	email varchar NULL,
	"password" varchar NULL,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);