package config

import (
	"github.com/spf13/viper"
)

// Config sdk for getting key values from settings file
type Config struct {
	viper *viper.Viper
}

// NewConfig initializes configuration from settings file
func NewConfig(defaults map[string]interface{}) *Config {
	config := &Config{viper: viper.New()}

	config.viper.SetDefault(dbHost, "localhost")
	config.viper.SetDefault(dbName, "gorm")
	config.viper.SetDefault(dbPort, 3306)
	config.viper.SetDefault(dbUser, "root")
	config.viper.SetDefault(dbPwd, "Pass@123")
	
	for key, value := range defaults {
		config.viper.SetDefault(key, value)
	}

	config.viper.AutomaticEnv()


	return config
}

// GetString return string value set for a given key
func (config *Config) GetString(key string) string {
	return config.viper.GetString(key)
}

// GetInt return int value set for the given key
func (config *Config) GetInt(key string) int {
	return config.viper.GetInt(key)
}

// Set sets the value for the given key
func (config *Config) Set(key string, value interface{}) {
	config.viper.Set(key, value)
}
