package app

import (
	"github.com/serenite11/market/services/order-service/config"
	api_grpc "github.com/serenite11/market/services/order-service/internal/api/grpc"
	"github.com/serenite11/market/services/order-service/internal/clients"
	order_usecase "github.com/serenite11/market/services/order-service/internal/domain/order/usecase"
	"github.com/serenite11/market/services/order-service/internal/factory"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		clients.NewModule(),
		factory.NewModule(),
		order_usecase.NewModule(),
		api_grpc.NewModule(),
		fx.Provide(
			config.New,
			logger.New,
		),
	)
}
