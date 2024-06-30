package delivery_grpc

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	grpc *grpc.Server
	cfg  *Config
	zap  *zap.Logger
}

func New(cfg *Config, zap *zap.Logger) *Server {
	var opts []grpc.ServerOption

	if cfg.TLSCredentialsPath != "" {
	}

	opts = append(opts, grpc.ChainUnaryInterceptor())

	return &Server{
		cfg:  cfg,
		grpc: grpc.NewServer(opts...),
		zap:  zap,
	}
}

func (s *Server) Start(ctx context.Context) error {
	reflection.Register(s.grpc)

	s.grpc.RegisterService(s.grpc)
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
