package server

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
)

const shutdownTimeout = 10 * time.Second
const httpAddr = ":8080"

func NewHTTPServer(logger *slog.Logger) (*http.Server, net.Listener) {
	mux := http.NewServeMux()
	// TODO: routing

	listener, err := net.Listen("tcp", httpAddr)
	if err != nil {
		logger.Error("HTTP server listen error", slog.Any("error", err))
		os.Exit(51)
	}

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	return server, listener
}

func StartHTTP(logger *slog.Logger, server *http.Server, listener net.Listener, started chan<- struct{}) {
	logger.Info("HTTP server starting...", slog.String("addr", server.Addr))
	close(started)
	if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
		logger.Error("HTTP server error", slog.Any("error", err))
		os.Exit(52)
	}
}

func StopHTTP(logger *slog.Logger, server *http.Server) {
	logger.Info("HTTP server stopping...")
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), shutdownTimeout)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("HTTP server shutdown error", slog.Any("error", err))

		logger.Info("HTTP server shutdown timed out, forcing close...")
		closingError := server.Close()
		if closingError != nil {
			logger.Error("HTTP server close error", slog.Any("error", closingError))
		}
	}
	logger.Info("HTTP server stopped")
}
