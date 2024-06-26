package config

import (
	"fmt"
	"go.uber.org/config"
	"go.uber.org/fx"
)

type Config struct {
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
