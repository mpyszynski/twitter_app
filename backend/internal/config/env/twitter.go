package env

import "github.com/spf13/viper"

type TwitterAuth struct {
	AppID string `mapstructure:"app_id"`
	AppKey string `mapstructure:"app_key"`
	AppSecret string `mapstructure:"app_secret"`
	AppToken string `mapstructure:"app_token"`
}

func getTwitterConfig() (*TwitterAuth, error) {
	var config *TwitterAuth
	for _, envVar := range []string{
		"app_id",
		"app_key",
		"app_secret",
		"app_token",
	} {
			if err := viper.BindEnv(envVar); err != nil {
				return nil, err
			}
	}
	viper.AutomaticEnv()
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}