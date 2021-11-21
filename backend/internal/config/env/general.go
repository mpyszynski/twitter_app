package env

import "github.com/spf13/viper"

// General holds general config for application
type General struct {
	Port string `mapstructure:"port"`
}

func getGeneralConfig() (*General, error) {
	var config *General
	for _, envVar := range []string{
		"port",
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