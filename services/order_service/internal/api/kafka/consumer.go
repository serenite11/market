package kafka

import (
	"context"
	"github.com/serenite11/market/services/order-service/pkg/kafka/consumer"
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type ConsumeFunc = func(ctx context.Context, msg *sarama.ConsumerMessage) error

type server struct {
	cfg      *kafka_consumer.Config
	consumer sarama.Consumer
	log      *zap.Logger
}

func New(cfg *kafka_consumer.Config, log *zap.Logger) *server {
	return &server{
		cfg: cfg,
		log: log,
	}
}

func (c *server) Init(_ context.Context) error {
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

func (c *server) Register() {
	c.consumer.Topics()
}

func (c *server) Listen(ctx context.Context, cb ConsumeFunc) error {
	partitions, err := c.consumer.ConsumePartition(c.cfg.Topic, 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}
	for v := range partitions.Messages() {
		if err = cb(ctx, v); err != nil {
			c.log.Error("Consumer error", zap.Error(err),
				zap.String("topic", c.cfg.Topic),
				zap.Any("consume_message", v),
			)
		}
	}

	return nil
}

func (c *server) Stop(_ context.Context) error {
	return c.consumer.Close()
}
