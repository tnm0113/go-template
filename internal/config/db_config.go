package config

type MongoDB struct {
	Host     string
	Port     int
	DbName   string
	Username string
	Password string
	Replica  string
}
