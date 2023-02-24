package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type ServiceConfig struct {
	DbConfig       MongoConfig    `mapstructure:"mongo"`
	GrpcConfig     GrpcConfig     `mapstructure:"grpc"`
	HttpConfig     HttpConfig     `mapstructure:"http"`
	LoggerConfig   LoggerConfig   `mapstructure:"logger"`
	RabbitmqConfig RabbitMQConfig `mapstructure:"rabbitmq"`
	OtherConfig    OtherConfig    `mapstructure:"other"`
}

func LoadConfig(path string) (cfg ServiceConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	setDefaultValue()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal().Msg("Load config fail")
		return
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal().Msg("Error when mapping config")
	}
	return
}

func setDefaultValue() {
	viper.SetDefault("other.environment", "development")

	viper.SetDefault("mongo.host", "127.0.0.1")
	viper.SetDefault("mongo.port", 27017)
	viper.SetDefault("mongo.name", "mongo")
	viper.SetDefault("mongo.user", "mongo")
	viper.SetDefault("mongo.pass", "mongo")

	viper.SetDefault("http.port", 8080)
	viper.SetDefault("http.enable_recover_middleware", true)
	viper.SetDefault("http.enable_cors_middleware", true)

	viper.SetDefault("grpc.port", 8081)

	viper.SetDefault("rabbitmq.username", "admin")
	viper.SetDefault("rabbitmq.password", "admin")
	viper.SetDefault("rabbitmq.vhost", "/")
	viper.SetDefault("rabbitmq.schema", "amqp")
	viper.SetDefault("rabbitmq.reconnect_max_attempt", "100")
	viper.SetDefault("rabbitmq.reconnect_interval", "5")
	viper.SetDefault("rabbitmq.channel_timeout", "5")

	viper.SetDefault("other.default_lang", "en")
	viper.SetDefault("other.bundle_dir_abs", "./web/i18n")
	viper.SetDefault("other.tracing_host", "127.0.0.1")
	viper.SetDefault("other.tracing_port", 9411)
}
