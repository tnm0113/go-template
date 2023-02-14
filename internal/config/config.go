package config

import "github.com/spf13/viper"

type ServiceConfig struct {
	Database MongoDB
	Host     string `mapstructure:"DB_HOST"`
}

func LoadConfig(path string) (cfg ServiceConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	viper.Unmarshal(&cfg)

	return
}
