package usecase

import (
	"context"
	"dynapgen/repository/github"
	"dynapgen/repository/nsq"
	"dynapgen/repository/redis"
	"dynapgen/util/cmd"
	"dynapgen/util/file"
	"dynapgen/util/log"
	"fmt"
	"strings"
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

	keystoreUrl, err := uc.github.Upload(ctx, github.UploadFileParam{
		FilePathLocal:       param.KeystorePath,
		FilePathRemote:      file.GenerateUniqueGithubFilePath(file.GithubFolderKeystore, file.GetFilenameFromPath(param.KeystorePath)),
		ReplaceIfNameExists: true,
	})
	if err != nil {
		log.Error(err, "uc.githb.Upload() got error - BuildApp", param)
		return err
	}

	// clean up local file
	cleanupCommand := fmt.Sprintf("rm %s", param.KeystorePath)
	if err := cmd.ExecCommandWithContext(ctx, cleanupCommand); err != nil {
		log.Error(err, "error cleaning up updated app icon - BuildApp", cleanupCommand)
		return err
	}

	if err = uc.cache.SetBuildAppStatus(ctx, redis.UpdateBuildStatusParam{
		AppID: param.AppID,
		BuildStatus: redis.BuildStatus{
			Status: redis.BuildStatusEnumPending,
		},
	}); err != nil {
		log.Error(err, "error initiating build app - BuildApp", nil)
		return err
	}

	buildParam := nsq.BuildAppParam{
		AppID:          param.AppID,
		AppName:        app.Name,
		AppVersionCode: param.VersionCode,
		AppVersionName: param.VersionName,
		TemplateType:   template.Type,
		TemplateName:   getRepositoryName(template.RepositoryURL),
		KeystoreUrl:    keystoreUrl,
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
