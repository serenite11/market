package app

import (
	api_grpc "github.com/serenite11/market/services/order-service/internal/api/grpc"
	order_usecase "github.com/serenite11/market/services/order-service/internal/domain/order/usecase"
	"github.com/serenite11/market/services/order-service/internal/factory"
	logger_zap "github.com/serenite11/market/services/order-service/pkg/logger"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		factory.NewModule(),
		order_usecase.NewModule(),
		api_grpc.NewModule(),
		fx.Provide(
			logger_zap.New(),
		),
	)
}
