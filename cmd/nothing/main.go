package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	slog.Info("Application starting..")
	slog.Info("Application started")

	<-gracefulStop
	slog.Info("Application stopping..")
	slog.Info("Application stopped")
}
