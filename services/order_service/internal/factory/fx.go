package factory

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"factory",
		fx.Provide(),
		fx.Invoke(),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("factory")
		}),
	)
}
