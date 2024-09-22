package api

import (
	api_grpc "github.com/serenite11/market/services/order-service/internal/api/grpc"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module("api",
		api_grpc.NewModule(),
		fx.Decorate(func(zap *zap.Logger) *zap.Logger {
			return zap.Named("api")
		}),
	)
}
