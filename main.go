package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/xeventa/base-service/app"
)

func main() {
	container, err := app.New()
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    container.Config.HTTPPort,
		Handler: container.Handler,
	}

	// Run server in a goroutine
	go func() {
		container.Logger.Info().Str("addr", srv.Addr).Msg("starting HTTP server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			container.Logger.Fatal().Err(err).Msg("http server error")
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	container.Logger.Info().Msg("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	if container.DB != nil {
		_ = container.DB.Close()
	}
}
