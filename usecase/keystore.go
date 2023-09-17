package usecase

import (
	"context"
	"dynapgen/repository/nsq"
	"dynapgen/repository/redis"
	"dynapgen/util/log"
)

func (uc *Usecase) BuildKeystore(ctx context.Context, param BuildKeystoreParam) error {
	err := uc.mq.PublishBuildKeystore(ctx, nsq.BuildKeystoreParam(param))
	if err != nil {
		log.Error(err, "error publishing build keystore message", param)
		return err
	}

	return nil
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
