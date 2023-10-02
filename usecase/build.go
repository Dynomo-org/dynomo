package usecase

import (
	"context"
	"dynapgen/repository/db"
	"dynapgen/repository/nsq"
	"dynapgen/repository/redis"
	"dynapgen/util/log"
	"strings"

	"github.com/rs/xid"
)

func (uc *Usecase) BuildApp(ctx context.Context, param BuildAppParam) error {
	app, err := uc.db.GetApp(ctx, param.AppID)
	if err != nil {
		log.Error(err, "uc.db.GetApp() got error - BuildApp")
		return err
	}

	if app.ID == "" {
		return nil
	}

	template, err := uc.GetTemplateByID(ctx, app.TemplateID)
	if err != nil {
		log.Error(err, "uc.GetTemplateByID() got error - BuildApp")
		return err
	}

	if template.ID == "" || template.RepositoryURL == "" {
		return nil
	}

	keystore, err := uc.db.GetKeystoreByID(ctx, param.KeystoreID)
	if err != nil {
		log.Error(err, "error getting keystore - BuildApp", param)
		return err
	}

	// generate build ID
	buildID := xid.New().String()
	if err = uc.cache.SetBuildAppStatus(ctx, redis.UpdateBuildStatusParam{
		BuildID: buildID,
		BuildStatus: redis.BuildStatus{
			Status: redis.BuildStatusEnumPending,
		},
	}); err != nil {
		log.Error(err, "error initiating build app - BuildApp", nil)
		return err
	}

	buildParam := nsq.BuildAppParam{
		BuildID:        buildID,
		AppName:        app.Name,
		AppVersionCode: app.VersionCode,
		AppVersionName: app.Version,
		TemplateType:   template.Type,
		TemplateName:   getRepositoryName(template.RepositoryURL),
		KeystoreUrl:    keystore.DownloadURL,
	}

	if err = uc.mq.PublishBuildApp(ctx, buildParam); err != nil {
		log.Error(err, "error publishing build app message - BuildApp", buildParam)
		return err
	}

	return nil
}

func (uc *Usecase) GetBuildAppStatus(ctx context.Context, appID string) (BuildStatus, error) {
	result, err := uc.cache.GetBuildAppStatus(ctx, appID)
	if err != nil {
		log.Error(err, "error getting build app status")
		return BuildStatus{}, err
	}

	return BuildStatus{
		Status:       BuildStatusEnum(result.Status),
		URL:          result.URL,
		ErrorMessage: result.ErrorMessage,
	}, nil
}

func (uc *Usecase) SetBuildAppStatus(ctx context.Context, param UpdateBuildStatusParam) error {
	if err := uc.cache.SetBuildAppStatus(ctx, redis.UpdateBuildStatusParam{
		BuildID: param.BuildID,
		BuildStatus: redis.BuildStatus{
			Status:       redis.BuildStatusEnum(param.BuildStatus.Status),
			URL:          param.BuildStatus.URL,
			ErrorMessage: param.BuildStatus.ErrorMessage,
		},
	}); err != nil {
		log.Error(err, "error updating build status to cache - SetBuildAppStatus", param)
		return err
	}

	if err := uc.db.UpsertAppArtifact(ctx, db.AppArtifact{
		ID:          param.BuildID,
		BuildStatus: int(param.BuildStatus.Status),
		DownloadURL: param.BuildStatus.URL,
	}); err != nil {
		log.Error(err, "error updating build status to db - SetBuildAppStatus", param)
		return err
	}

	return nil
}

func getRepositoryName(url string) string {
	segments := strings.Split(url, "/")
	return segments[len(segments)-1]
}
