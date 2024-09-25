package config

import (
	"fmt"
	"go.uber.org/config"
	"go.uber.org/fx"
)

type Config struct {
	Grpc  GrpcServer `yaml:"grpc"`
	Kafka Kafka      `yaml:"kafka"`
}

type GrpcServer struct {
	TransportName string `yaml:"transport_name"`
	Port          int    `yaml:"port"`
}

type Producers struct{}

type Kafka struct {
	Addrs []string `yaml:"addrs"`
	Topic string   `yaml:"topic"`
}

type ResultConfig struct {
	fx.Out
	Config *Config
}

func New() (ResultConfig, error) {
	loader, err := config.NewYAML(config.File("config/config.yaml"))
	if err != nil {
		return ResultConfig{}, fmt.Errorf("%w", err)
	}
	cfg := Config{}

	if err = loader.Get("app").Populate(&cfg); err != nil {
		return ResultConfig{}, err
	}

	return ResultConfig{
		Config: &cfg,
	}, nil
}
