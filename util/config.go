package util

import "github.com/spf13/viper"

// Config stores all configuration of the app
// The values are read by viper from a config file or env. variables
type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	HttpServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
}

// LoadConfig reads the configuration from file or env. variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// Overwrite values with values from config files
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
