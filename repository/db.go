package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const (
	collectionMstApp = "mst_app"

	AdTypeAdmob       = 1
	AdTypeMAN         = 2
	AdTypeAppLovinMax = 3

	AdmobOpenAdTestID         = "ca-app-pub-3940256099942544/3419835294"
	AdmobBannerAdTestID       = "ca-app-pub-3940256099942544/6300978111"
	AdmobInterstitialAdTestID = "ca-app-pub-3940256099942544/1033173712"
	AdmobRewardAdTestID       = "ca-app-pub-3940256099942544/5224354917"
	AdmobNativeAdTestID       = "ca-app-pub-3940256099942544/2247696110"
)

func (r *Repository) GetAllAppFromDB(ctx context.Context) ([]App, error) {
	var result map[string]interface{}
	err := r.db.NewRef(collectionMstApp).Get(ctx, &result)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}

	slice := make([]interface{}, 0, len(result))
	for _, app := range result {
		slice = append(slice, app)
	}
	sliceStr, _ := json.Marshal(slice)

	var data []App
	err = json.Unmarshal([]byte(sliceStr), &data)
	return data, err
}

func (r *Repository) GetAppFromDB(ctx context.Context, appID string) (App, error) {
	var result App
	err := r.db.NewRef(collectionMstApp).Child(appID).Get(ctx, &result)
	return result, err
}

func (r *Repository) InsertAppToDB(ctx context.Context, master App) error {
	_, err := r.db.NewRef(collectionMstApp).Push(ctx, master)
	return err
}

func (r *Repository) InsertNewAppToDB(ctx context.Context, request NewAppRequest) error {
	appID := uuid.NewString()
	master := App{
		AppID:          appID,
		AppName:        request.AppName,
		AppPackageName: request.PackageName,
		VersionCode:    1,
		VersionName:    "1.0",
		IconURL:        "https://raw.githubusercontent.com/Dynapgen/master-storage-1/main/assets/default-icon.png",
		CreatedAt:      time.Now(),
		AdsConfig: AdsConfig{
			PrimaryAd: Ad{
				AdType:           AdTypeAdmob,
				OpenAdID:         AdmobOpenAdTestID,
				BannerAdID:       AdmobBannerAdTestID,
				InterstitialAdID: AdmobInterstitialAdTestID,
				RewardAdID:       AdmobRewardAdTestID,
				NativeAdID:       AdmobNativeAdTestID,
			},
		},
		AppConfig: AppConfig{
			Strings: AppString{
				SetAsWallpaper:      "Set As Wallpaper",
				SetWallpaperHome:    "Home Wallpaper",
				SetWallpaperLock:    "Lockscreen Wallpaper",
				WallpaperBoth:       "Home + Lock Screen",
				Cancel:              "Cancel",
				SuccessSetWallpaper: "Set Wallpaper Success",
				ExitPromptMessage:   "Do You Want to Exit?",
				NoConnectionMessage: "No Connection",
				PrivacyPolicyText:   "",
			},
			Style: AppStyle{
				ColorPrimary:          "#FFBB86FC",
				ColorPrimaryVariant:   "#FF3700B3",
				ColorSecondary:        "#FF03DAC5",
				ColorSecondaryVariant: "#FF018786",
				ColorOnPrimary:        "#FFFFFFFF",
				ColorOnSecondary:      "#FF000000",
			},
		},
	}
	err := r.db.NewRef(collectionMstApp).Child(appID).Set(ctx, master)
	return err
}

func (r *Repository) UpdateAppOnDB(ctx context.Context, app App) error {
	return r.db.NewRef(collectionMstApp).Update(ctx,
		map[string]interface{}{
			app.AppID: app,
		})
}

func (r *Repository) DeleteAppOnDB(ctx context.Context, appID string) error {
	return r.db.NewRef(collectionMstApp).Child(appID).Delete(ctx)
}
