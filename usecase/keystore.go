package usecase

import (
	"context"
	"dynapgen/repository/db"
	"dynapgen/repository/nsq"
	"dynapgen/repository/redis"
	"dynapgen/util/crypto"
	"dynapgen/util/log"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/rs/xid"
)

func (uc *Usecase) BuildKeystore(ctx context.Context, param BuildKeystoreParam) error {
	// generate buildID
	buildID := xid.New().String()

	metadata := keystoreMetadata{
		Alias:         param.Alias,
		FullName:      param.FullName,
		Organization:  param.Organization,
		Country:       param.Country,
		KeyPassword:   param.KeyPassword,
		StorePassword: param.StorePassword,
	}

	marshalledMetadata, err := json.Marshal(metadata)
	if err != nil {
		log.Error(err, "error marshalling keystore metadata", metadata)
		return err
	}

	encryptedMetadata, err := crypto.EncryptAES(string(marshalledMetadata))
	if err != nil {
		log.Error(err, "error encrypting keystore metadata", string(marshalledMetadata))
		return err
	}

	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
		errs  []error
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := uc.db.UpsertKeystore(ctx, db.Keystore{
			ID:          buildID,
			OwnerID:     param.OwnerID,
			Name:        param.KeystoreName,
			Alias:       param.Alias,
			Metadata:    encryptedMetadata,
			BuildStatus: int(BuildStatusEnumPending),
			CreatedAt:   time.Now(),
		}); err != nil {
			log.Error(err, "error inserting keystore to db", param)
			mutex.Lock()
			errs = append(errs, err)
			mutex.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := uc.cache.SetBuildKeystoreStatus(ctx, redis.UpdateBuildStatusParam{
			BuildID: buildID,
			BuildStatus: redis.BuildStatus{
				Status: redis.BuildStatusEnumPending,
			},
		}); err != nil {
			log.Error(err, "error initiating build keystore")
			mutex.Lock()
			errs = append(errs, err)
			mutex.Unlock()
		}
	}()
	wg.Wait()

	if len(errs) != 0 {
		err = errors.New("build keystore failed")
		log.Error(err, "error updating keystore status", map[string]interface{}{
			"param":  param,
			"errors": errs,
		})
		return err
	}

	if err := uc.mq.PublishBuildKeystore(ctx, nsq.BuildKeystoreParam{
		BuildID:       buildID,
		FullName:      param.FullName,
		Organization:  param.Organization,
		Country:       param.Country,
		Alias:         param.Alias,
		KeyPassword:   param.KeyPassword,
		StorePassword: param.StorePassword,
	}); err != nil {
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

func (uc *Usecase) GetKeystoreList(ctx context.Context, param GetKeystoreListParam) (GetKeystoreListResponse, error) {
	keystores, err := uc.db.GetKeystoresByOwnerID(ctx, param.OwnerID)
	if err != nil {
		log.Error(err, "error getting keystore list", param)
		return GetKeystoreListResponse{}, err
	}

	result := make([]Keystore, len(keystores))
	for index, keystore := range keystores {
		result[index] = convertKeystoreFromDB(keystore)
	}

	return buildKeystoreListResponse(result, param), nil
}

func (uc *Usecase) SetBuildKeystoreStatus(ctx context.Context, param UpdateBuildStatusParam) error {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
		errs  []error
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := uc.db.UpsertKeystore(ctx, db.Keystore{
			ID:          param.BuildID,
			BuildStatus: int(param.BuildStatus.Status),
			DownloadURL: param.BuildStatus.URL,
		}); err != nil {
			log.Error(err, "error updating keystore status in DB", param)
			mutex.Lock()
			errs = append(errs, err)
			mutex.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		if err := uc.cache.SetBuildKeystoreStatus(ctx, redis.UpdateBuildStatusParam{
			BuildID: param.BuildID,
			BuildStatus: redis.BuildStatus{
				Status:       redis.BuildStatusEnum(param.BuildStatus.Status),
				URL:          param.BuildStatus.URL,
				ErrorMessage: param.BuildStatus.ErrorMessage,
			},
		}); err != nil {
			log.Error(err, "error updating keystore status in cache", param)
			mutex.Lock()
			errs = append(errs, err)
			mutex.Unlock()
		}
	}()
	wg.Wait()

	if len(errs) != 0 {
		log.Error(nil, "error updating keystore status", map[string]interface{}{
			"param":  param,
			"errors": errs,
		})
		return errors.New("update keystore status failed")
	}

	return nil
}

func buildKeystoreListResponse(keystores []Keystore, param GetKeystoreListParam) GetKeystoreListResponse {
	if len(keystores) == 0 {
		return GetKeystoreListResponse{
			Keystores: []Keystore{},
		}
	}

	return GetKeystoreListResponse{
		Keystores: keystores,
		TotalPage: keystores[0].Total / param.PerPage,
		Page:      param.Page,
	}
}
