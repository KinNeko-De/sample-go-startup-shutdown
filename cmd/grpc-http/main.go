package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/kinneko-de/sample-go-startup-shutdown/internal/app/operation/health"
	"github.com/kinneko-de/sample-go-startup-shutdown/internal/app/server"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	slog.Info("Application starting..")

	grpcServer, grpcHealthServer, grpcListener := server.NewGRPCServer(logger)
	httpServer, httpListener := server.NewHTTPServer(logger)
	grpcServerStarted := make(chan struct{})
	httpServerStarted := make(chan struct{})
	go server.StartGRPC(logger, grpcServer, grpcListener, grpcServerStarted)
	go server.StartHTTP(logger, httpServer, httpListener, httpServerStarted)

	<-grpcServerStarted
	<-httpServerStarted
	health.Ready(grpcHealthServer)

	slog.Info("Application started")

	<-gracefulStop
	slog.Info("Application stopping..")

	health.NotReady(grpcHealthServer)

	server.StopHTTP(logger, httpServer)
	server.StopGRPC(logger, grpcServer)

	slog.Info("Application stopped")
}
