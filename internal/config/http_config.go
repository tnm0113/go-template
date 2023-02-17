package config

type HttpConfig struct {
	//HTTP config
	HttpHost                string `mapstructure:"host"`
	HttpPort                int    `mapstructure:"port"`
	EnableRecoverMiddleware bool   `mapstructure:"enable_recover_middleware"`
	EnableCORSMiddleware    bool   `mapstructure:"enable_cors_middleware"`
	EchoDebug               bool   `mapstructure:"ECHO_DEBUG"`
}
