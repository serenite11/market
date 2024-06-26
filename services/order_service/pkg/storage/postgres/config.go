package storage_postgres

import (
	"fmt"

	"go.uber.org/config"
)

type Config struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	Database       string `yaml:"database"`
	SslMode        string `yaml:"ssl_mode"`
	MigrationsPath string `yaml:"migrations_path"`
}

func NewConfig(provider config.Provider) (*Config, error) {
	var cfg Config
	err := provider.Get("postgres").Populate(&cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot parse cfg %w", err)
	}
	return &cfg, err
}
