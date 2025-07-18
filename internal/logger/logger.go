package logger

import (
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	return log
}