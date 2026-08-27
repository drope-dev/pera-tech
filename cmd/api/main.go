package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/drope-dev/pera-tech/internal/catalog"
	"github.com/drope-dev/pera-tech/internal/config"
	"github.com/drope-dev/pera-tech/internal/platform/database"
	"github.com/drope-dev/pera-tech/internal/platform/httpapi"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.Error("invalid application configuration")
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	db, err := database.Connect(ctx, cfg.DatabaseURL, cfg.DatabaseMaxConns)
	if err != nil {
		logger.Error("database connection failed")
		os.Exit(1)
	}
	defer db.Close()
	products := catalog.NewService(catalog.NewPostgresRepository(db))

	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           httpapi.NewHandler(db, products, cfg.AdminAPIToken),
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	serverErr := make(chan error, 1)
	go func() {
		logger.Info("http server started", "address", cfg.HTTPAddr)
		serverErr <- server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Error("http server shutdown failed")
			os.Exit(1)
		}
	case err := <-serverErr:
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Error("http server stopped unexpectedly")
			os.Exit(1)
		}
	}
}
