package order_usecase

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("order_usecase",
		fx.Provide(
			fx.Annotate(New, fx.As(new(UseCase))),
		),
		fx.Invoke(func(lc fx.Lifecycle, uc UseCase) {
			lc.Append(fx.Hook{
				OnStart: func(context.Context) error {
					return nil
				},
				OnStop: func(context.Context) error {
					return nil
				},
			})
		}),
		fx.Decorate(func(zap *zap.Logger) *zap.Logger {
			return zap.Named("order_usecase")
		}),
	)
}
