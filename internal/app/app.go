package app

import (
	"fmt"
	"github/nergilz/taskGetRate/internal/server"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	logger     *slog.Logger
	port       string
	grpcServer *grpc.Server
}

func New(log *slog.Logger, port string, rateService server.IGetRates) *App {
	grpcServer := grpc.NewServer()

	server.Register(grpcServer, rateService)

	return &App{
		logger:     log,
		port:       port,
		grpcServer: grpcServer,
	}
}

func (app *App) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", app.port))
	if err != nil {
		panic(fmt.Errorf("app.Run: %w", err))
	}

	app.logger.Info("run grpc server", slog.String("addr:", listener.Addr().String()))

	if err := app.grpcServer.Serve(listener); err != nil {
		panic(fmt.Errorf("app.Run: %w", err))
	}
}

func (app *App) Stop() {
	app.logger.Info("stop grpc server", slog.String("port:", app.port))

	app.grpcServer.GracefulStop()
}
