package clients

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("clients",
		fx.Provide(),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("clients")
		}),
	)
}
