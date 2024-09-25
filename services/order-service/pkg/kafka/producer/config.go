package kafka_producer

import "go.uber.org/config"

type Config struct {
	Addrs []string
}

func NewConfig(provider config.Provider) (*Config, error) {
	var cfg Config
	if err := provider.Get("kafka_producer").Populate(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
