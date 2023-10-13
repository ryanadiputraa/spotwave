package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	*Server
}

type Server struct {
	Port string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return
}
