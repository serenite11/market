package kafka

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("kafka",
		fx.Provide(
			fx.Annotate(),
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("kafka")
		}),
	)
}
