package api_grpc

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewModule() fx.Option {
	return fx.Module(
		"grpc",
		fx.Provide(
			New,
		),
		fx.Invoke(func(lc fx.Lifecycle, server *grpc.Server) {
			lc.Append(fx.Hook{
				OnStart: nil,
				OnStop:  nil,
			})
		}),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("grpc")
		}),
	)
}
