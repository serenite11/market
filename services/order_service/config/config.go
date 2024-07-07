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

type ResultConfig struct {
	fx.Out
	Provider config.Provider
}

func New() (ResultConfig, error) {
	loader, err := config.NewYAML(config.File("config/config.yaml"))
	if err != nil {
		return ResultConfig{}, fmt.Errorf("%w", err)
	}
	return ResultConfig{
		Provider: loader,
	}, nil
}
