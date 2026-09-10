package main

import (
	"log/slog"
	"os"

	"github.com/borntow1n/pastebin-mini/internal/config"
	"github.com/borntow1n/pastebin-mini/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	logger := setupLogger(cfg.Env)
	db, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		logger.Error("failed to initialisation database")
	}

}

func setupLogger(env string) *slog.Logger {
	switch env {
	case "local":
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "dev":
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		panic("unknown environment: " + env)
	}
}
