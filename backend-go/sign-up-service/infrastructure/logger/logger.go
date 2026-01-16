package infrastructure_logger

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

var AppLogger *slog.Logger

// InitLogger configures the global logger based on .env settings

func InitLogger() {
	logLevel := os.Getenv("LOG_LEVEL")

	var level slog.Level
	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		level = slog.LevelDebug
	case "WARN":
		level = slog.LevelWarn
	case "ERROR":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	logDir := "logs"

	// Use our new Daily Writer instead of a standard os.File
	dailyWriter := NewDailyFileWriter(logDir)

	// MultiWriter now sends to Console + our smart DailyWriter
	multiWriter := io.MultiWriter(os.Stdout, dailyWriter)

	opts := &slog.HandlerOptions{
		Level: level,
	}

	baseHandler := slog.NewJSONHandler(multiWriter, opts)

	finalHandler := &ContextHandler{Handler: baseHandler}

	slog.SetDefault(slog.New(finalHandler))
}
