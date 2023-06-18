package usecase

import (
	"context"
	"dynapgen/utils/cmd"
	"dynapgen/utils/log"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (uc *Usecase) BuildApp(ctx context.Context, param BuildAppParam) (interface{}, error) {
	go uc.buildAppAsync(ctx, param)

	return nil, nil
}

func (uc *Usecase) buildAppAsync(ctx context.Context, param BuildAppParam) {
	var (
		appFolderPath string
		err           error
	)
	defer func() {
		// os.RemoveAll(param.KeystorePath)
		// os.RemoveAll(appFolderPath)
	}()

	app, err := uc.db.GetApp(ctx, param.AppID)
	if err != nil {
		log.Error(nil, err, "uc.db.GetApp() got error - buildAppAsync")
		return
	}

	if app.ID == "" {
		return
	}

	template, err := uc.GetTemplateByID(ctx, app.TemplateID)
	if err != nil {
		log.Error(nil, err, "uc.GetTemplateByID() got error - buildAppAsync")
		return
	}

	if template.ID == "" || template.RepositoryURL == "" {
		return
	}

	keystoreNameSegments := strings.Split(param.KeystorePath, "/")
	keystoreFileName := keystoreNameSegments[len(keystoreNameSegments)-1]
	keystoreFolderPath := filepath.Join("./garage", param.AppID)
	err = os.MkdirAll(keystoreFolderPath, 0755)
	if err != nil {
		log.Error(keystoreFolderPath, err, "cmd.ExecCommandWithContext() got error - buildAppAsync")
		return
	}

	keystorePath := filepath.Join(keystoreFolderPath, keystoreFileName)
	mvKeystoreCommand := fmt.Sprintf("mv %s %s", param.KeystorePath, keystorePath)
	err = cmd.ExecCommandWithContext(ctx, mvKeystoreCommand)
	if err != nil {
		log.Error(mvKeystoreCommand, err, "cmd.ExecCommandWithContext() got error - buildAppAsync")
		return
	}

	repoFolderPath := filepath.Join("./templates", getRepositoryName(template.RepositoryURL))
	appFolderPath = filepath.Join("./garage", param.AppID+"/"+getRepositoryName(template.RepositoryURL))
	if _, err := os.Stat(repoFolderPath); os.IsNotExist(err) {
		cloneCommand := fmt.Sprintf("cd ./templates && git clone %s", template.RepositoryURL)
		err = cmd.ExecCommandWithContext(ctx, cloneCommand)
		if err != nil {
			log.Error(cloneCommand, err, "cmd.ExecCommandWithContext() got error - buildAppAsync")
			return
		}
	}

	cpProjectCommand := fmt.Sprintf("cp -R %s %s", repoFolderPath, appFolderPath)
	err = cmd.ExecCommandWithContext(ctx, cpProjectCommand)
	if err != nil {
		log.Error(cpProjectCommand, err, "cmd.ExecCommandWithContext() got error - buildAppAsync")
		return
	}

	// TODO: Algorithm to inject app data, and icon
	// TODO: Execute build command
	// TODO: Upload artifact to github

	log.Info("AMAN")
}

func getRepositoryName(url string) string {
	segments := strings.Split(url, "/")
	return segments[len(segments)-1]
}
