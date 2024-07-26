package kafka_producer

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("producer",
		fx.Provide(
			NewConfig,
			fx.Annotate(New, fx.As(new(Producer))),
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("producer")
		}),
	)
}
