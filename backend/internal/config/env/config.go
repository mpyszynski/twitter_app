package env

import (
	"github.com/spf13/viper"
)



type Config struct {
	TwitterAuth *TwitterAuth
	AppID string `mapstructure:"app_id"`
	AppKey string `mapstructure:"app_key"`
	AppSecret string `mapstructure:"app_secret"`
	AppToken string `mapstructure:"app_token"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("secrets")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	twitterAuth, err := getTwitterConfig()
	if err != nil {
		return nil, err
	}
	config := Config{TwitterAuth: twitterAuth}
	return &config, nil
}

