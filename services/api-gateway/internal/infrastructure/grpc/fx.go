package grpc

import (
	"github.com/serenite11/market/services/api-gateway/internal/infrastructure/grpc/clients/orderservice"
	"go.uber.org/fx"
)

var Module = fx.Module("grpc",
	fx.Provide(
		orderservice.NewClient,
	),
	fx.Invoke(
		func(lc fx.Lifecycle, c *orderservice.Client) {
			lc.Append(fx.Hook{
				OnStart: c.OnStart,
				OnStop:  c.OnStop,
			})
		},
	),
)
