package config

import (
	"fmt"
	"go.uber.org/config"
	"go.uber.org/fx"
	delivery_grpc "order-service/internal/domain/order/delivery/grpc"
)

type Config struct {
	Grpc delivery_grpc.Config `yaml:"grpc"`
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
