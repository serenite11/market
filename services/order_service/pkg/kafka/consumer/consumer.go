package kafka_consumer

import (
	"context"
	"github.com/IBM/sarama"
	"time"
)

type ConsumeFunc = func(ctx context.Context, msg *sarama.ConsumerMessage) error

type Consumer interface {
	Start(ctx context.Context) error
	StartConsume(ctx context.Context, cb ConsumeFunc) error
	Stop(ctx context.Context) error
}

type client struct {
	cfg *Config
	sarama.Consumer
}

func New(cfg *Config) Consumer {
	return &client{
		cfg: cfg,
	}
}

func (c *client) Start(_ context.Context) error {
	saramaConfig := sarama.NewConfig()

	saramaConfig.Version = sarama.V1_0_0_0
	saramaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	saramaConfig.Consumer.Return.Errors = true
	saramaConfig.Net.DialTimeout = time.Second * 3

	client, err := sarama.NewConsumer(c.cfg.Addrs, saramaConfig)
	if err != nil {
		return err
	}
	c.Consumer = client
	return nil
}

func (c *client) StartConsume(ctx context.Context, cb ConsumeFunc) error {
	return nil
}

func (c *client) Stop(_ context.Context) error {
	return c.Close()
}
