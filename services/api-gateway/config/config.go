package config

type Config struct {
}

type OrderService struct {
	GrpcClientConfig `yaml:",inline"`
}

type GrpcClientConfig struct {
	DSN         string `yaml:"DSN"`
	TLSRequired bool   `yaml:"TLSRequired"`
}
