package domain

import (
	order_usecase "github.com/serenite11/market/services/order-service/internal/domain/order/usecase"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("domain",
		order_usecase.NewModule(),
		fx.Decorate(func(zap *zap.Logger) *zap.Logger {
			return zap.Named("domain")
		}),
	)
}
