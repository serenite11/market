package kafka_consumer

import (
	"github.com/serenite11/market/services/order-service/internal/api/kafka"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("consumer",
		fx.Provide(
			NewConfig,
			fx.Annotate(kafka.New, fx.As(new(kafka.Consumer))),
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("consumer")
		}),
	)
}
