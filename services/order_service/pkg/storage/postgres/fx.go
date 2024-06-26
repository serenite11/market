package storage_postgres

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("postgres",
		fx.Provide(
			NewConfig,
			fx.Annotate(New, fx.As(new(Postgres))),
		),
		fx.Invoke(func(lc fx.Lifecycle, p Postgres) {
			lc.Append(fx.Hook{
				OnStart: p.Connect,
				OnStop:  p.Close,
			})
		}),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("postgres")
		}),
	)
}
