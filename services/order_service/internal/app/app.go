package app

import (
	"github.com/serenite11/market/services/order-service/config"
	"github.com/serenite11/market/services/order-service/internal/api"
	"github.com/serenite11/market/services/order-service/internal/domain"
	"github.com/serenite11/market/services/order-service/internal/infra/factory"
	"github.com/serenite11/market/services/order-service/pkg/logger"
	"go.uber.org/fx"
)

func CreateApp() fx.Option {
	return fx.Options(
		factory.NewModule(),
		domain.NewModule(),
		api.NewModule(),
		fx.Provide(
			config.New,
			logger.New,
		),
	)
}
