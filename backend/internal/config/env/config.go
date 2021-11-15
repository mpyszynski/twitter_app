package env

import (
	"github.com/spf13/viper"
)



type Config struct {
	TwitterAuth *TwitterAuth
}

func LoadConfig(path, fileName string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(fileName)
	viper.AutomaticEnv()
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

