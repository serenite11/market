package api_grpc

import (
	"context"
	"github.com/serenite11/market/proto/services/order_service_v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	order_usecase "order-service/internal/domain/order/usecase"
)

type Server struct {
	grpc *grpc.Server
	cfg  *Config
	zap  *zap.Logger
	uc   order_usecase.UseCase
}

func New(cfg *Config, zap *zap.Logger, uc order_usecase.UseCase) *Server {
	var opts []grpc.ServerOption

	if cfg.TLSCredentialsPath != "" {
	}

	opts = append(opts, grpc.ChainUnaryInterceptor())

	return &Server{
		cfg:  cfg,
		grpc: grpc.NewServer(opts...),
		zap:  zap,
		uc:   uc,
	}
}

func (s *Server) Start(ctx context.Context) error {
	reflection.Register(s.grpc)

	grpcHandler := NewOrderHandler(s.uc)

	order_service_v1.RegisterOrderServiceServer(s.grpc, grpcHandler)

	go func() {
		if err := s.run(ctx); err != nil {
			s.zap.Error("failed to start server", zap.Error(err))
		}
	}()

	return nil
}

func (s *Server) Stop(_ context.Context) error {
	s.grpc.GracefulStop()
	return nil
}

func (s *Server) run(_ context.Context) error {
	listener, err := net.Listen("tcp", ":"+s.cfg.Port)
	if err != nil {
		return err
	}

	return s.grpc.Serve(listener)
}
