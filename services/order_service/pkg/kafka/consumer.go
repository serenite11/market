package kafka

import (
	"context"
	kafka_consumer "github.com/serenite11/market/services/order-service/internal/api/kafka"
)

type Consumer interface {
	Init(ctx context.Context) error
	StartConsume(ctx context.Context, cb kafka_consumer.ConsumeFunc) error
	Stop(ctx context.Context) error
}
