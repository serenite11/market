package factory

import (
	storage_postgres "github.com/serenite11/market/services/order-service/pkg/storage/postgres"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"factory",
		storage_postgres.NewModule(),
		fx.Provide(
			New,
		),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("factory")
		}),
	)
}
