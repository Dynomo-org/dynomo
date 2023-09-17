package env

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
)

const (
	KeyGithubUsername = "DNM_GH_USERNAME"
	KeyGithubToken    = "DNM_GH_TOKEN"

	keyENV = "DNM_ENV"
)

var env string

func InitConfig() error {
	// singleton initialization
	// ensure no double initialization
	if env != "" {
		return nil
	}

	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	env = viper.GetString(keyENV)
	if env == "" {
		return errors.New("env not found")
	}

	return nil
}

func IsDevelopment() bool {
	return strings.ToLower(env) == "development"
}
