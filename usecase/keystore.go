package usecase

import (
	"context"
	"dynapgen/repository/nsq"
	"dynapgen/repository/redis"
	"dynapgen/util/log"
)

func (uc *Usecase) BuildKeystore(ctx context.Context, param BuildKeystoreParam) error {
	if err := uc.cache.SetBuildKeystoreStatus(ctx, redis.UpdateBuildStatusParam{
		AppID: param.AppID,
		BuildStatus: redis.BuildStatus{
			Status: redis.BuildStatusEnumPending,
		},
	}); err != nil {
		log.Error(err, "error initiating build keystore")
		return err
	}

	if err := uc.mq.PublishBuildKeystore(ctx, nsq.BuildKeystoreParam(param)); err != nil {
		log.Error(err, "error publishing build keystore message", param)
		return err
	}

	return nil
}

func (uc *Usecase) GetBuildKeystoreStatus(ctx context.Context, appID string) (BuildStatus, error) {
	result, err := uc.cache.GetBuildKeystoreStatus(ctx, appID)
	if err != nil {
		log.Error(err, "error getting build keystore status")
		return BuildStatus{}, err
	}

	return BuildStatus{
		Status:       BuildStatusEnum(result.Status),
		URL:          result.URL,
		ErrorMessage: result.ErrorMessage,
	}, nil
}

func (uc *Usecase) SetBuildKeystoreStatus(ctx context.Context, param UpdateBuildStatusParam) error {
	return uc.cache.SetBuildKeystoreStatus(ctx, redis.UpdateBuildStatusParam{
		AppID: param.AppID,
		BuildStatus: redis.BuildStatus{
			Status:       redis.BuildStatusEnum(param.BuildStatus.Status),
			URL:          param.BuildStatus.URL,
			ErrorMessage: param.BuildStatus.ErrorMessage,
		},
	})
}
