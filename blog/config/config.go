package config

import "github.com/spf13/viper"

type Config struct {
	Port        string `mapstructure:"port"`
	DatabaseUrl string `mapstructure:"database_url"`
	PublicKey   string `mapstructure:"public_key"`
}

// Return a config from app.env file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
