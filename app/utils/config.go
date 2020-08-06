package utils

import "github.com/spf13/viper"

// Config struct
type Config struct {
	DBName     string
	DBUser     string
	DBPassword string
	Secret     string
	Port       string
	TestDB     string
}

var config *Config

// GetConfig func
func GetConfig() *Config {
	return config
}

// LoadConfigVars -- load configuration
func LoadConfig() (*Config, error) {
	// Set environment file
	viper.SetConfigFile(".env")

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		return new(Config), err
	}

	config = &Config{
		DBName:     viper.GetString("DB_NAME"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		Secret:     viper.GetString("SECRET"),
		Port:       viper.GetString("PORT"),
		TestDB:     viper.GetString("TEST_DB"),
	}

	return config, nil
}
