// config/config.go
package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config struct holds the application configuration
type Config struct {
	DB     DBConfig   `yaml:"db"`
	Server ServerConf `yaml:"server"`
}

// DBConfig holds the database configuration
type DBConfig struct {
	Driver   string `yaml:"driver"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

// ServerConf holds the server configuration
type ServerConf struct {
	Port int `yaml:"port"`
}

// LoadConfig loads the configuration from a YAML file
func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
