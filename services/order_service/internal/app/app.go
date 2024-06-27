package app

import (
	"go.uber.org/fx"
	logger_zap "order-service/pkg/logger"
)

func New() *fx.App {
	return fx.New(

		fx.Provide(
			logger_zap.New(),
		),
	)
}
