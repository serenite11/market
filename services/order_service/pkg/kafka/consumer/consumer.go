package kafka_consumer

import (
	"context"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"time"
)

type ConsumeFunc = func(ctx context.Context, msg *sarama.ConsumerMessage) error

type Consumer interface {
	Init(ctx context.Context) error
	StartConsume(ctx context.Context, cb ConsumeFunc) error
	Stop(ctx context.Context) error
}

type client struct {
	cfg      *Config
	consumer sarama.Consumer
	log      *zap.Logger
}

func New(cfg *Config, log *zap.Logger) Consumer {
	return &client{
		cfg: cfg,
		log: log,
	}
}

func (c *client) Init(_ context.Context) error {
	saramaConfig := sarama.NewConfig()

	saramaConfig.Version = sarama.V1_0_0_0
	saramaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	saramaConfig.Consumer.Return.Errors = true
	saramaConfig.Net.DialTimeout = time.Second * 3

	client, err := sarama.NewConsumer(c.cfg.Addrs, saramaConfig)
	if err != nil {
		return err
	}
	c.consumer = client
	return nil
}

func (c *client) StartConsume(ctx context.Context, cb ConsumeFunc) error {
	partitions, err := c.consumer.ConsumePartition(c.cfg.Topic, 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}
	for v := range partitions.Messages() {
		if err = cb(context.Background(), v); err != nil {
			c.log.Error("Consumer error", zap.Error(err),
				zap.String("topic", c.cfg.Topic),
				zap.Any("consume_message", v),
			)
		}
	}

	return nil
}

func (c *client) Stop(_ context.Context) error {
	return c.consumer.Close()
}
