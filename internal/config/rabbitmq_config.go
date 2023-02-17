package config

type RabbitMQConfig struct {
	Schema               string `mapstructure:"schema"`
	Username             string `mapstructure:"username"`
	Password             string `mapstructure:"password"`
	Host                 string `mapstructure:"host"`
	Port                 int    `mapstructure:"port"`
	Vhost                string `mapstructure:"vhost"`
	ChannelNotifyTimeout int    `mapstructure:"channel_timeout"`
	ConnectionName       string `mapstructure:"connection_name"`
	ReconnectMaxAttempt  int    `mapstructure:"reconnect_max_attempt"`
	ReconnectInterval    int    `mapstructure:"reconnect_interval"`
}
