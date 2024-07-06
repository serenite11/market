package api_grpc

import "go.uber.org/fx"

func NewModule() fx.Option {
	return fx.Module(
		"grpc",
	)
}
