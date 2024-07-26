package kafka

import (
	kafka_producer "github.com/serenite11/market/services/order-service/pkg/kafka/producer"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("kafka",
		kafka_producer.NewModule(),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("kafka")
		}),
	)
}
