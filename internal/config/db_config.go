package config

type MongoDB struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	DbName   string `mapstructure:"DB_NAME"`
	Username string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASS"`
	Replica  string `mapstructure:"DB_REPLICA"`
}
