package usecase

import (
	"context"
	"dynapgen/repository/github"
	"dynapgen/repository/redis"
	"dynapgen/utils/cmd"
	"dynapgen/utils/log"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func (uc *Usecase) GetGenerateKeystoreStatus(ctx context.Context, appID string) (Keystore, error) {
	keystore, err := uc.cache.GetKeystore(ctx, appID)
	if err != nil {
		log.Error(map[string]interface{}{"app_id": appID}, err, "uc.cache.GetKeystore() got error - GetGenerateKeystoreStatus")
		return Keystore{}, err
	}

	return convertKeystoreFromRepo(keystore), nil
}

func (uc *Usecase) GenerateKeystore(ctx context.Context, param GenerateStoreParam) error {
	err := uc.cache.InvalidateKeystore(ctx, param.AppID)
	if err != nil {
		log.Error(param, err, "uc.cache.InvalidateKeystore() got error - generateKeystoreAsync")
		return err
	}

	keystore := redis.Keystore{
		Status: redis.BuildStatusPending,
	}

	err = uc.cache.InsertKeystore(ctx, param.AppID, keystore)
	if err != nil {
		log.Error(param, err, "uc.repo.InsertKeystore() got error - GenerateKeystore")
		return err
	}

	ctxAsync, cancel := context.WithTimeout(ctx, time.Duration(60)*time.Second)
	go func() {
		defer cancel()
		uc.generateKeystoreAsync(ctxAsync, param)
	}()

	return nil
}

func (uc *Usecase) generateKeystoreAsync(ctx context.Context, param GenerateStoreParam) {
	var (
		keystore   redis.Keystore
		folderPath string
		err        error
	)
	defer func() {
		if folderPath != "" && err == nil {
			err = os.RemoveAll(folderPath)
		}
		if err != nil {
			keystore = redis.Keystore{
				Status:       redis.BuildStatusFail,
				ErrorMessage: err.Error(),
			}
			err := uc.cache.InsertKeystore(ctx, param.AppID, keystore)
			if err != nil {
				log.Error(param, err, "uc.repo.InsertKeystore() got error - generateKeystoreAsync")
			}
		}
	}()

	keystore, err = uc.cache.GetKeystore(ctx, param.AppID)
	if err != nil {
		log.Error(param, err, "uc.repo.GetKeystore() got error - generateKeystoreAsync")
		return
	}

	keystore.Status = redis.BuildStatusInProgress
	err = uc.cache.InsertKeystore(ctx, param.AppID, keystore)
	if err != nil {
		log.Error(param, err, "uc.repo.InsertKeystore() got error - generateKeystoreAsync")
		return
	}

	folderPath = filepath.Join("./assets/", param.AppID)
	err = os.RemoveAll(folderPath)
	if err != nil {
		log.Error(folderPath, err, "os.RemoveAll() got error - generateKeystoreAsync")
		return
	}

	err = os.MkdirAll(folderPath, 0755)
	if err != nil {
		log.Error(folderPath, err, "os.Mkdir() got error - generateKeystoreAsync")
		return
	}

	filePath := "./assets/" + param.AppID + "/key.keystore"
	dName := fmt.Sprintf("CN=%s,O=%s,C=%s", param.FullName, param.Organization, param.Country)

	fullCommand := fmt.Sprintf("keytool -genkey -v -keystore %s -storepass %s -alias %s -keypass %s -keyalg RSA -keysize 2048 -validity 10000 -noprompt -dname %s", filePath, param.StorePassword, param.Alias, param.KeyPassword, dName)
	err = cmd.ExecCommandWithContext(ctx, fullCommand)
	if err != nil {
		log.Error(param, err, "cmd.ExecCommandWithContext() got error - generateKeystoreAsync")
		return
	}

	uploadGithubParam := github.UploadFileParam{
		FilePathLocal:         filePath,
		FileName:              "key.keystore",
		DestinationFolderPath: param.AppID + "/",
		ReplaceIfNameExists:   true,
	}

	fileURL, err := uc.github.Upload(ctx, uploadGithubParam)
	if err != nil {
		log.Error(uploadGithubParam, err, "uc.repo.Upload() got error - generateKeystoreAsync")
		return
	}

	keystore.Status = redis.BuildStatusSuccess
	keystore.URL = fileURL
	err = uc.cache.InsertKeystore(ctx, param.AppID, keystore)
	if err != nil {
		log.Error(param, err, "uc.cache.InsertKeystore() got error - generateKeystoreAsync")
		return
	}

	log.Info("[Keystore] Generated: " + param.AppID)
}
