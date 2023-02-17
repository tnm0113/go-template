package config

import (
	"github.com/spf13/viper"
)

type ServiceConfig struct {
	DbConfig       MongoConfig
	GrpcConfig     GrpcConfig
	HttpConfig     HttpConfig
	LoggerConfig   LoggerConfig
	RabbitmqConfig RabbitMQConfig
	OtherConfig    OtherConfig
}

func LoadConfig(path string) (cfg ServiceConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	setDefaultValue()

	err = viper.ReadInConfig()

	viper.UnmarshalKey("mongo", &cfg.DbConfig)
	viper.UnmarshalKey("grpc", &cfg.GrpcConfig)
	viper.UnmarshalKey("http", &cfg.HttpConfig)
	viper.UnmarshalKey("log", &cfg.LoggerConfig)
	viper.UnmarshalKey("rabbitmq", &cfg.RabbitmqConfig)
	viper.UnmarshalKey("other", &cfg.OtherConfig)

	return
}

func setDefaultValue() {
	viper.SetDefault("other.environment", "development")

	viper.SetDefault("mongo.host", "127.0.0.1")
	viper.SetDefault("mongo.port", 27017)
	viper.SetDefault("mongo.name", "mongo")
	viper.SetDefault("mongo.user", "mongo")
	viper.SetDefault("mongo.pass", "mongo")
}
