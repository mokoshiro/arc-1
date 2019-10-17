package rpc

import (
	"fmt"
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opencensus.io/plugin/ocgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Bo0km4n/arc/pkg/rpc/interceptor"
)

type Service interface {
	Embed(server *grpc.Server, logger *zap.Logger)
}

type Server struct {
	port       int
	services   []Service
	grpcServer *grpc.Server
	logger     *zap.Logger
}

type Option func(*Server)

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func WithService(service Service) Option {
	return func(s *Server) {
		s.services = append(s.services, service)
	}
}

func WithLogger(logger *zap.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func NewServer(service Service, opts ...Option) *Server {
	logger, _ := zap.NewProduction()
	s := &Server{
		logger: logger,
	}
	for _, opt := range opts {
		opt(s)
	}
	s.logger = s.logger.Named("rpc-server")
	if service == nil {
		s.logger.Fatal("service must not be nil")
	}
	s.services = append(s.services, service)
	return s
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		s.logger.Fatal("failed to listen", zap.Error(err))
	}

	opts := []grpc.ServerOption{
		grpc.StatsHandler(&ocgrpc.ServerHandler{}),
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				interceptor.StreamServerLoggingIntercepter(s.logger.Named("grpc.server.stream")),
				interceptor.RequestValidationStreamServerInterceptor(),
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				interceptor.ServerLoggingIntercepter(s.logger.Named("grpc.server.unary")),
				interceptor.RequestValidationUnaryServerInterceptor(),
			),
		),
	}
	s.grpcServer = grpc.NewServer(opts...)
	for _, service := range s.services {
		service.Embed(s.grpcServer, s.logger)
	}

	s.logger.Info(fmt.Sprintf("running on %s", lis.Addr().String()))
	defer s.logger.Sync()
	reflection.Register(s.grpcServer)
	err = s.grpcServer.Serve(lis)
	if err != nil && err != grpc.ErrServerStopped {
		s.logger.Fatal("failed to serve", zap.Error(err))
	}
}

func (s *Server) Stop(timeout time.Duration) {
	ch := make(chan struct{})
	go func() {
		s.logger.Info("gracefulStop is running")
		s.grpcServer.GracefulStop()
		close(ch)
	}()
	select {
	case <-ch:
		s.logger.Info("gracefulStop completed before timing out")
	case <-time.After(timeout):
		s.logger.Info("force server to stop")
		s.grpcServer.Stop()
	}
}
