package usecase

import (
	"context"
	"dynapgen/repository"
	"dynapgen/utils/log"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func (uc *Usecase) GetGenerateKeystoreStatus(ctx context.Context, appID string) (Keystore, error) {
	keystore, err := uc.repo.GetKeystoreFromCache(ctx, appID)
	if err != nil {
		log.Error(map[string]interface{}{"app_id": appID}, err, "uc.repo.GetKeystoreFromCache() got error - GetGenerateKeystoreStatus")
		return Keystore{}, err
	}

	return convertKeystoreFromRepo(keystore), nil
}

func (uc *Usecase) GenerateKeystore(ctx context.Context, param GenerateStoreParam) error {
	err := uc.repo.InvalidateKeystoreOnCache(ctx, param.AppID)
	if err != nil {
		log.Error(param, err, "uc.repo.GetKeystoreFromCache() got error - generateKeystoreAsync")
		return err
	}

	keystore := repository.Keystore{
		Status: repository.BuildStatusPending,
	}

	err = uc.repo.StoreKeystoreToCache(ctx, param.AppID, keystore)
	if err != nil {
		log.Error(param, err, "uc.repo.StoreKeystoreToCache() got error - GenerateKeystore")
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
		keystore   repository.Keystore
		folderPath string
		err        error
	)
	defer func() {
		if folderPath != "" {
			err = os.RemoveAll(folderPath)
		}

		if err != nil {
			keystore = repository.Keystore{
				Status:       repository.BuildStatusFail,
				ErrorMessage: err.Error(),
			}
			err := uc.repo.StoreKeystoreToCache(ctx, param.AppID, keystore)
			if err != nil {
				log.Error(param, err, "uc.repo.StoreKeystoreToCache() got error - generateKeystoreAsync")
			}
		}
	}()

	keystore, err = uc.repo.GetKeystoreFromCache(ctx, param.AppID)
	if err != nil {
		log.Error(param, err, "uc.repo.GetKeystoreFromCache() got error - generateKeystoreAsync")
		return
	}

	keystore.Status = repository.BuildStatusInProgress
	err = uc.repo.StoreKeystoreToCache(ctx, param.AppID, keystore)
	if err != nil {
		log.Error(param, err, "uc.repo.StoreKeystoreToCache() got error - generateKeystoreAsync")
		return
	}

	folderPath = filepath.Join("./assets/", param.AppID)
	err = os.RemoveAll(folderPath)
	if err != nil {
		log.Error(folderPath, err, "os.RemoveAll() got error - generateKeystoreAsync")
		return
	}

	err = os.Mkdir(folderPath, 0755)
	if err != nil {
		log.Error(folderPath, err, "os.Mkdir() got error - generateKeystoreAsync")
		return
	}

	filePath := "./assets/" + param.AppID + "/key.keystore"
	dName := fmt.Sprintf("CN=%s,O=%s,C=%s", param.FullName, param.Organization, param.Country)
	cmd := exec.CommandContext(ctx, "keytool", "-genkey", "-v", "-keystore", filePath, "-storepass", param.StorePassword, "-alias", param.Alias, "-keypass", param.KeyPassword, "-keyalg", "RSA", "-keysize", "2048", "-validity", "10000", "-noprompt", "-dname", dName)

	_, err = cmd.Output()
	if err != nil {
		log.Error(param, err, "exec.Command() got error - generateKeystoreAsync")
		return
	}

	uploadGithubParam := repository.UploadFileParam{
		FilePathLocal:         filePath,
		FileName:              "key.keystore",
		DestinationFolderPath: param.AppID + "/",
		ReplaceIfNameExists:   true,
	}

	fileURL, err := uc.repo.UploadToGithub(ctx, uploadGithubParam)
	if err != nil {
		log.Error(uploadGithubParam, err, "uc.repo.UploadToGithub() got error - generateKeystoreAsync")
		return
	}

	keystore.Status = repository.BuildStatusSuccess
	keystore.URL = fileURL
	err = uc.repo.StoreKeystoreToCache(ctx, param.AppID, keystore)
	if err != nil {
		log.Error(param, err, "uc.repo.StoreKeystoreToCache() got error - generateKeystoreAsync")
		return
	}

	log.Info("[Keystore] Generated: " + param.AppID)
}
