// configloader.go
package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

var dbConfig DBConfig
var redisConfig RedisConfig

func init() {
	// Set the name of the config file (without extension)
	viper.SetConfigName("config")
	// Set the path to look for the config file
	viper.AddConfigPath(".")
	// Set the file extension
	viper.SetConfigType("yaml")

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return
	}

	// Unmarshal the database configuration into the struct
	err = viper.UnmarshalKey("database", &dbConfig)
	if err != nil {
		fmt.Printf("Error unmarshaling database config: %s\n", err)
		return
	}

	// Unmarshal the Redis configuration into the struct
	err = viper.UnmarshalKey("redis", &redisConfig)
	if err != nil {
		fmt.Printf("Error unmarshaling redis config: %s\n", err)
		return
	}
}

func GetDBConfig() DBConfig {
	return dbConfig
}

func GetRedisConfig() RedisConfig {
	return redisConfig
}
