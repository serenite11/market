package kafka_consumer

import "go.uber.org/config"

type Config struct {
	Addrs []string `yaml:"addrs"`
	Topic string   `yaml:"topic"`
	Group string   `yaml:"group"`
}

func NewConfig(provider config.Provider) (*Config, error) {
	cfg := &Config{}
	if err := provider.Get("kafka").Populate(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
