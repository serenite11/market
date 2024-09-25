package api_grpc

import (
	"go.uber.org/config"
)

type Config struct {
	Port               string `yaml:"port"`
	TLSCredentialsPath string `yaml:"tls_credentials_path"`
	TLSEnabled         bool   `yaml:"tls_enabled"`
}

func NewConfig(provider config.Provider) (*Config, error) {
	cfg := Config{}
	if err := provider.Get("grpc").Populate(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
