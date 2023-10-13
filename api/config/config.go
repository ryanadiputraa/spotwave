package config

import (
	"os"

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
	env := getEnv()

	viper.SetConfigName(env)
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return
}

func getEnv() (env string) {
	if len(os.Args) <= 1 {
		return "development"
	}

	env = os.Args[1]
	switch env {
	case "prod":
		env = "production"
	case "stage":
		env = "staging"
	default:
		env = "development"
	}
	return
}
