package api_grpc

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"grpc",
		fx.Provide(
			NewConfig,
			New,
		),
		fx.Invoke(func(lc fx.Lifecycle, server *Server) {
			lc.Append(fx.Hook{
				OnStart: server.Start,
				OnStop:  server.Stop,
			})
		}),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("grpc")
		}),
	)
}
