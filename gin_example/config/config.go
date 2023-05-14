package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string `json:"port"`
}

func (s *Config) NewConfig() *Config {
	return &Config{}
}

func (s *Config) ServerConfig() *Config {
	viper.SetConfigFile("../.env")
	viper.ReadInConfig()

	s.Port = viper.GetString("PORT")

	return s
}
