package usecase

import (
	"context"
	"dynapgen/repository/nsq"
	"dynapgen/repository/redis"
	"dynapgen/util/log"
	"strings"
)

func (uc *Usecase) BuildApp(ctx context.Context, param BuildAppParam) error {
	app, err := uc.db.GetApp(ctx, param.AppID)
	if err != nil {
		log.Error(err, "uc.db.GetApp() got error - buildAppAsync")
		return err
	}

	if app.ID == "" {
		return nil
	}

	template, err := uc.GetTemplateByID(ctx, app.TemplateID)
	if err != nil {
		log.Error(err, "uc.GetTemplateByID() got error - buildAppAsync")
		return err
	}

	if template.ID == "" || template.RepositoryURL == "" {
		return nil
	}

	buildParam := nsq.BuildAppParam{
		AppID:          param.AppID,
		AppName:        app.Name,
		AppVersionCode: param.VersionCode,
		AppVersionName: param.VersionName,
		TemplateType:   template.Type,
		TemplateName:   getRepositoryName(template.RepositoryURL),
	}

	err = uc.mq.PublishBuildApp(ctx, buildParam)
	if err != nil {
		log.Error(err, "error publishing build app message", buildParam)
		return err
	}

	return nil
}

func (uc *Usecase) SetBuildAppStatus(ctx context.Context, param UpdateBuildStatusParam) error {
	return uc.cache.SetBuildAppStatus(ctx, redis.UpdateBuildStatusParam{
		AppID: param.AppID,
		BuildStatus: redis.BuildStatus{
			Status:       redis.BuildStatusEnum(param.BuildStatus.Status),
			URL:          param.BuildStatus.URL,
			ErrorMessage: param.BuildStatus.ErrorMessage,
		},
	})
}

func getRepositoryName(url string) string {
	segments := strings.Split(url, "/")
	return segments[len(segments)-1]
}
