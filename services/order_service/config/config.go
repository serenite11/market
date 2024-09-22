package config

import (
	"fmt"
	api_grpc "github.com/serenite11/market/services/order-service/internal/api/grpc"
	"go.uber.org/config"
	"go.uber.org/fx"
)

type Config struct {
	Grpc api_grpc.Config `yaml:"grpc"`
}

type GrpcServer struct {
	TransportName string `yaml:"transport_name"`
	Port          int    `yaml:"port"`
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
