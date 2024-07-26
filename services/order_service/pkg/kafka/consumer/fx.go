package kafka_consumer

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("consumer",
		fx.Provide(
			NewConfig,
			fx.Annotate(New, fx.As(new(Consumer))),
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("consumer")
		}),
	)
}
