package app

import (
	"go.uber.org/fx"
	"order-service/internal/domain/order/usecase"
	logger_zap "order-service/pkg/logger"
)

func New() *fx.App {
	return fx.New(
		usecase.NewModule(),
		fx.Provide(
			logger_zap.New(),
		),
	)
}
