package main

import (
	"github/nergilz/taskGetRate/internal/app"
	"github/nergilz/taskGetRate/internal/config"
	"github/nergilz/taskGetRate/internal/service"
	"github/nergilz/taskGetRate/internal/storage"
	"github/nergilz/taskGetRate/internal/transport"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// run app
func main() {
	cfg := config.Load()

	logger := SetupLogger(cfg.Env)

	logger.Info("init config")

	appTransport := transport.New(cfg.GrinexBaseUrl)
	appStorage := storage.New()
	appService := service.New(logger, appStorage, appTransport)
	appRates := app.New(logger, cfg.GrpcCfg.Port, appService)

	go func() {
		appRates.Run()
	}()

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGTERM, syscall.SIGINT)

	<-shutdownCh

	appRates.Stop()
	logger.Info("appliation stoped")
}

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "dev":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return log
}
