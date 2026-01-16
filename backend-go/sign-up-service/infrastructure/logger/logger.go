package infrastructure_logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"
)

var AppLogger *slog.Logger

// InitLogger configures the global logger based on .env settings

func InitLogger() {
	logLevel := os.Getenv("LOG_LEVE")

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

	// 1. Create the filename based on the current date: yyyy-mm-dd_logs.txt
	currentTime := time.Now().Format("2006-01-02")
	fileName := fmt.Sprintf("%s_logs.txt", currentTime)

	// 2. Open the file (Append mode, Create if not exists, Read/Write permissions)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
	}

	// 3. Create a MultiWriter to log to BOTH the Console and the File
	multiWriter := io.MultiWriter(os.Stdout, file)

	opts := &slog.HandlerOptions{
		Level: level,
	}

	baseHandler := slog.NewJSONHandler(multiWriter, opts)

	finalHandler := &ContextHandler{Handler: baseHandler}

	slog.SetDefault(slog.New(finalHandler))
}
