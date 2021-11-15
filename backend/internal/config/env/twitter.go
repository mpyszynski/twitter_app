package env

import "github.com/spf13/viper"

type TwitterAuth struct {
	AppID string `mapstructure:"app_id"`
	ApiKey string `mapstructure:"api_key"`
	ApiSecret string `mapstructure:"api_secret"`
	ApiToken string `mapstructure:"api_token"`
	ApiTokenSecret string `mapstructure:"api_token_secret"`
}

func getTwitterConfig() (*TwitterAuth, error) {
	var config *TwitterAuth
	for _, envVar := range []string{
		"app_id",
		"api_key",
		"api_secret",
		"api_token",
		"api_token_secret",
	} {
			if err := viper.BindEnv(envVar); err != nil {
				return nil, err
			}
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}