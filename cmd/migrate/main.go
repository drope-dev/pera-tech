package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/drope-dev/pera-tech/internal/config"
	"github.com/drope-dev/pera-tech/internal/platform/database"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.Error("invalid application configuration")
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db, err := database.Connect(ctx, cfg.DatabaseURL, cfg.DatabaseMaxConns)
	if err != nil {
		logger.Error("database connection failed")
		os.Exit(1)
	}
	defer db.Close()

	if err := database.ApplyMigrations(ctx, db); err != nil {
		logger.Error("database migration failed")
		os.Exit(1)
	}

	logger.Info("database migrations applied")
}
