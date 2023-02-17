package config

type GrpcConfig struct {
	//GRPC config
	GrpcPort int    `mapstructure:"port"`
	GrpcHost string `mapstructure:"host"`
}
