package server

import (
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"
	grpcHealth "google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/kinneko-de/sample-go-startup-shutdown/internal/app/operation/health"
)

const grpcAddr = ":9090"

func NewGRPCServer(logger *slog.Logger) (*grpc.Server, *grpcHealth.Server, net.Listener) {
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logger.Error("failed to create gRPC listener", slog.Any("error", err))
		os.Exit(41)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	healthServer := grpcHealth.NewServer()
	grpc_health_v1.RegisterHealthServer(server, healthServer)
	health.Initialize(healthServer)

	return server, healthServer, listener
}

func StartGRPC(logger *slog.Logger, server *grpc.Server, listener net.Listener, started chan<- struct{}) {
	logger.Info("gRPC server starting", slog.Any("addr", listener.Addr()))
	close(started)
	if err := server.Serve(listener); err != nil {
		logger.Error("gRPC server error", slog.Any("error", err))
		os.Exit(42)
	}
}

func StopGRPC(logger *slog.Logger, server *grpc.Server) {
	logger.Info("gRPC server stopping...")
	server.GracefulStop()
	logger.Info("gRPC server stopped")
}
