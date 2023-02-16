package config

import "github.com/spf13/viper"

type ServiceConfig struct {
	//DB config
	DBHost    string `mapstructure:"DB_HOST"`
	DBPort    int    `mapstructure:"DB_PORT"`
	DBName    string `mapstructure:"DB_NAME"`
	DBUser    string `mapstructure:"DB_USER"`
	DBPass    string `mapstructure:"DB_PASS"`
	DBReplica string `mapstructure:"DB_REPLICA"`

	//HTTP config
	HttpHost                string `mapstructure:"HTTP_HOST"`
	HttpPort                int    `mapstructure:"HTTP_PORT"`
	EnableRecoverMiddleware bool   `mapstructure:"ENABLE_RECOVER_MIDDLEWARE"`
	EnableCORSMiddleware    bool   `mapstructure:"ENABLE_CORS_MIDDLEWARE"`
	EchoDebug               bool   `mapstructure:"ECHO_DEBUG"`

	//GRPC config
	GrpcPort int `mapstructure:"GRPC_PORT"`
	GrpcHost int `mapstructure:"GRPC_HOST"`

	//Logger config
	LogLevel           int  `mapstructure:"LOG_LEVEL"`
	RequestLevel       int  `mapstructure:"REQUEST_LEVEL"`
	PrettyPrintConsole bool `mapstructure:"PRETTY_PRINT_CONSOLE"`

	//i18n
	DefaultLanguage string `mapstructure:"DEFAULT_LANGUAGE"`
	BundleDirAbs    string `mapstructure:"BUNDLE_DIR_ABS"`
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

func SetDefaultValue() {

}
