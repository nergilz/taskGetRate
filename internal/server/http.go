package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"sync"
)

type Handler struct {
	Logger   *slog.Logger
	ServeMux *http.ServeMux
}

func NewHandler(log *slog.Logger) Handler {
	mux := http.NewServeMux()

	h := Handler{
		Logger:   log,
		ServeMux: mux,
	}

	h.addRoutes(mux)

	return h
}

func (h Handler) addRoutes(mux *http.ServeMux) {
	mux.Handle("/healthcheck", h.handleHealthcheck())
}

func (h Handler) handleHealthcheck() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Welcome to the Home Page: OK!")

			h.Logger.Info("handle home page")
		},
	)
}

func Run(ctx context.Context, handler http.Handler) {
	log.Println("Server starting on :8080...")

	httpServer := &http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", "8080"),
		Handler: handler,
	}

	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()

		err := httpServer.Shutdown(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()

	wg.Wait()
}
