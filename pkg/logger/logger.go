package logger

import (
	"RalfDalfs/go-fiber/config"
	"log/slog"
	"os"
)

func NewLogger(config *config.LogConfig) *slog.Logger {
	slog.SetLogLoggerLevel(slog.Level(config.Level))
	var logger *slog.Logger
	if config.Format == "json" {
		logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
	return logger
}
