package config

type LoggerConfig struct {
	LogLevel           int  `mapstructure:"log_level"`
	RequestLevel       int  `mapstructure:"request_level"`
	PrettyPrintConsole bool `mapstructure:"pretty_print_console"`
}
