package env

import (
	"dynapgen/util/log"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	KeyGithubUsername = "DNM_GH_USERNAME"
	KeyGithubToken    = "DNM_GH_TOKEN"

	keyENV = "ENV"
)

var env string

func InitConfig() error {
	// singleton initialization
	// ensure no double initialization
	if env != "" {
		return nil
	}

	viper.AutomaticEnv()
	env = viper.GetString(keyENV)
	if env == "" {
		env = "development"
	}
	log.Info(fmt.Sprintf("running on env %s", env))

	viper.SetConfigFile(fmt.Sprintf("./config/config.%s.yaml", env))
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func IsDevelopment() bool {
	return strings.ToLower(env) == "development"
}
