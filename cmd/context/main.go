package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, stopListen := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopListen()

	slog.Info("Application starting..")

	var wg sync.WaitGroup
	wg.Add(1)
	go worker(ctx, &wg)

	slog.Info("Application started")

	<-ctx.Done()
	slog.Info("Application stopping..")
	wg.Wait() // wait for the worker to finish its cleanup
	slog.Info("Application stopped")
}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			slog.Info("Worker stopping..")
			shutdownWorker()
			slog.Info("Worker stopped")
			return
		case <-time.After(1 * time.Second):
			slog.Info("Working..")
		}
	}
}

func shutdownWorker() {
	time.Sleep(2 * time.Second)
}
