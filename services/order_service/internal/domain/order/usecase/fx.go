package order_usecase

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("order_usecase",
		fx.Provide(
			fx.Annotate(New, fx.As(new(UseCase))),
		),
		fx.Decorate(func(zap *zap.Logger) *zap.Logger {
			return zap.Named("order_usecase")
		}),
	)
}
