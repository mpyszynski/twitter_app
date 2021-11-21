package env

import (
	"github.com/spf13/viper"
)

// Config holds all configurations for application
type Config struct {
	TwitterAuth *TwitterAuth
	General     *General
}

// LoadConfig loads configuration
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
	generalConfig, err := getGeneralConfig()
	if err != nil {
		return nil, err
	}
	config := Config{TwitterAuth: twitterAuth, General: generalConfig}
	return &config, nil
}
