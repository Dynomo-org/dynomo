package usecase

import (
	"dynapgen/repository"
	"time"
)

type NewAppRequest struct {
	AppName     string
	PackageName string
}

type App struct {
	AppID             string
	AppName           string
	AppPackageName    string
	VersionCode       uint
	VersionName       string
	IconURL           string
	PrivacyPolicyLink string
	AdmobAppID        string
	AppLovinSDKKey    string
	AdsConfig         AdsConfig
	AppConfig         AppConfig
	Contents          []AppContent
	Categories        []AppCategory
	CreatedAt         time.Time
	UpdatedAt         *time.Time
}

type AdsConfig struct {
	EnableOpenAd         bool
	EnableBannerAd       bool
	EnableInterstitialAd bool
	EnableRewardAd       bool
	EnableNativeAd       bool

	PrimaryAd   Ad
	SecondaryAd Ad
	TertiaryAd  Ad

	InterstitialIntervalSecond int
	TestDevices                []string
}

type Ad struct {
	AdType           uint8
	OpenAdID         string
	BannerAdID       string
	InterstitialAdID string
	RewardAdID       string
	NativeAdID       string
}

type AppCategory struct {
	ID       string
	Category string
}

type AppContent struct {
	ID          string
	Title       string
	Description string
	Content     string
	CategoryID  string
}

type AppConfig struct {
	Strings AppString
	Style   AppStyle
}

type AppString struct {
	SetAsWallpaper      string
	SetWallpaperHome    string
	SetWallpaperLock    string
	WallpaperBoth       string
	Cancel              string
	SuccessSetWallpaper string
	ExitPromptMessage   string
	NoConnectionMessage string
	PrivacyPolicyText   string
}

type AppStyle struct {
	ColorPrimary          string
	ColorPrimaryVariant   string
	ColorOnPrimary        string
	ColorSecondary        string
	ColorSecondaryVariant string
	ColorOnSecondary      string
}

func convertAppFromRepo(app repository.App) App {
	contents := make([]AppContent, 0, len(app.Contents))
	for _, content := range app.Contents {
		contents = append(contents, AppContent(content))
	}

	categories := make([]AppCategory, 0, len(app.Categories))
	for _, category := range app.Categories {
		categories = append(categories, AppCategory(category))
	}

	return App{
		AppID:             app.AppID,
		AppName:           app.AppName,
		AppPackageName:    app.AppPackageName,
		VersionCode:       app.VersionCode,
		VersionName:       app.VersionName,
		IconURL:           app.IconURL,
		PrivacyPolicyLink: app.PrivacyPolicyLink,
		AdsConfig: AdsConfig{
			EnableOpenAd:               app.AdsConfig.EnableOpenAd,
			EnableBannerAd:             app.AdsConfig.EnableBannerAd,
			EnableInterstitialAd:       app.AdsConfig.EnableInterstitialAd,
			EnableRewardAd:             app.AdsConfig.EnableRewardAd,
			EnableNativeAd:             app.AdsConfig.EnableNativeAd,
			PrimaryAd:                  Ad(app.AdsConfig.PrimaryAd),
			SecondaryAd:                Ad(app.AdsConfig.SecondaryAd),
			TertiaryAd:                 Ad(app.AdsConfig.TertiaryAd),
			InterstitialIntervalSecond: app.AdsConfig.InterstitialIntervalSecond,
			TestDevices:                app.AdsConfig.TestDevices,
		},
		AppConfig: AppConfig{
			Strings: AppString(app.AppConfig.Strings),
			Style:   AppStyle(app.AppConfig.Style),
		},
		Contents:   contents,
		Categories: categories,
		CreatedAt:  app.CreatedAt,
		UpdatedAt:  app.UpdatedAt,
	}
}

func (app *App) convertAppToRepo() repository.App {
	contents := make([]repository.AppContent, 0, len(app.Contents))
	for _, content := range app.Contents {
		contents = append(contents, repository.AppContent(content))
	}

	categories := make([]repository.AppCategory, 0, len(app.Categories))
	for _, category := range app.Categories {
		categories = append(categories, repository.AppCategory(category))
	}

	return repository.App{
		AppID:             app.AppID,
		AppName:           app.AppName,
		AppPackageName:    app.AppPackageName,
		VersionCode:       app.VersionCode,
		VersionName:       app.VersionName,
		IconURL:           app.IconURL,
		PrivacyPolicyLink: app.PrivacyPolicyLink,
		AdsConfig: repository.AdsConfig{
			EnableOpenAd:               app.AdsConfig.EnableOpenAd,
			EnableBannerAd:             app.AdsConfig.EnableBannerAd,
			EnableInterstitialAd:       app.AdsConfig.EnableInterstitialAd,
			EnableRewardAd:             app.AdsConfig.EnableRewardAd,
			EnableNativeAd:             app.AdsConfig.EnableNativeAd,
			PrimaryAd:                  repository.Ad(app.AdsConfig.PrimaryAd),
			SecondaryAd:                repository.Ad(app.AdsConfig.SecondaryAd),
			TertiaryAd:                 repository.Ad(app.AdsConfig.TertiaryAd),
			InterstitialIntervalSecond: app.AdsConfig.InterstitialIntervalSecond,
			TestDevices:                app.AdsConfig.TestDevices,
		},
		AppConfig: repository.AppConfig{
			Strings: repository.AppString(app.AppConfig.Strings),
			Style:   repository.AppStyle(app.AppConfig.Style),
		},
		Contents:   contents,
		CreatedAt:  app.CreatedAt,
		UpdatedAt:  app.UpdatedAt,
		Categories: categories,
	}
}

// updateWith will update the app with the given input ONLY if the input attribute is not empty
func (app *App) updateWith(input App) {
	if input.AppName != "" {
		app.AppName = input.AppName
	}
	if input.AppPackageName != "" {
		app.AppPackageName = input.AppPackageName
	}
	if input.VersionCode != 0 {
		app.VersionCode = input.VersionCode
	}
	if input.VersionName != "" {
		app.VersionName = input.VersionName
	}
	if input.IconURL != "" {
		app.IconURL = input.IconURL
	}
	if input.PrivacyPolicyLink != "" {
		app.PrivacyPolicyLink = input.PrivacyPolicyLink
	}

	// style settings
	if input.AppConfig.Style.ColorPrimary != "" {
		app.AppConfig.Style.ColorPrimary = input.AppConfig.Style.ColorPrimary
	}
	if input.AppConfig.Style.ColorPrimaryVariant != "" {
		app.AppConfig.Style.ColorPrimaryVariant = input.AppConfig.Style.ColorPrimaryVariant
	}
	if input.AppConfig.Style.ColorOnPrimary != "" {
		app.AppConfig.Style.ColorOnPrimary = input.AppConfig.Style.ColorOnPrimary
	}
	if input.AppConfig.Style.ColorSecondary != "" {
		app.AppConfig.Style.ColorSecondary = input.AppConfig.Style.ColorSecondary
	}
	if input.AppConfig.Style.ColorSecondaryVariant != "" {
		app.AppConfig.Style.ColorSecondaryVariant = input.AppConfig.Style.ColorSecondaryVariant
	}
	if input.AppConfig.Style.ColorOnSecondary != "" {
		app.AppConfig.Style.ColorOnSecondary = input.AppConfig.Style.ColorOnSecondary
	}

	// string settings
	if input.AppConfig.Strings.SetAsWallpaper != "" {
		app.AppConfig.Strings.SetAsWallpaper = input.AppConfig.Strings.SetAsWallpaper
	}
	if input.AppConfig.Strings.SetWallpaperHome != "" {
		app.AppConfig.Strings.SetWallpaperHome = input.AppConfig.Strings.SetWallpaperHome
	}
	if input.AppConfig.Strings.SetWallpaperLock != "" {
		app.AppConfig.Strings.SetWallpaperLock = input.AppConfig.Strings.SetWallpaperLock
	}
	if input.AppConfig.Strings.WallpaperBoth != "" {
		app.AppConfig.Strings.WallpaperBoth = input.AppConfig.Strings.WallpaperBoth
	}
	if input.AppConfig.Strings.Cancel != "" {
		app.AppConfig.Strings.Cancel = input.AppConfig.Strings.Cancel
	}
	if input.AppConfig.Strings.SuccessSetWallpaper != "" {
		app.AppConfig.Strings.SuccessSetWallpaper = input.AppConfig.Strings.SuccessSetWallpaper
	}
	if input.AppConfig.Strings.ExitPromptMessage != "" {
		app.AppConfig.Strings.ExitPromptMessage = input.AppConfig.Strings.ExitPromptMessage
	}
	if input.AppConfig.Strings.NoConnectionMessage != "" {
		app.AppConfig.Strings.NoConnectionMessage = input.AppConfig.Strings.NoConnectionMessage
	}
	if input.AppConfig.Strings.PrivacyPolicyText != "" {
		app.AppConfig.Strings.PrivacyPolicyText = input.AppConfig.Strings.PrivacyPolicyText
	}
}

type GenerateStoreParam struct {
	AppID         string
	FullName      string
	Organization  string
	Country       string
	Alias         string
	KeyPassword   string
	StorePassword string
}

type Keystore struct {
	Status       uint8
	URL          string
	ErrorMessage string
}

func convertKeystoreFromRepo(keystore repository.Keystore) Keystore {
	return Keystore{
		Status:       uint8(keystore.Status),
		URL:          keystore.URL,
		ErrorMessage: keystore.ErrorMessage,
	}
}
