package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	*Server
	*Google
	*Spotify
	*RapidAPI
}

type Server struct {
	Port   string `mapstructure:"port"`
	WebURL string `mapstructure:"webUrl"`
}

type Google struct {
	APIKey string `mapstructure:"apiKey"`
}

type Spotify struct {
	ClientID     string `mapstructure:"clientId"`
	ClientSecret string `mapstructure:"clientSecret"`
	State        string `mapstructure:"state"`
	RedirectURI  string `mapstructure:"redirectUri"`
}

type RapidAPI struct {
	APIKey string `mapstructure:"apiKey"`
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
