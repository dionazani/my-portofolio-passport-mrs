package middleware

import (
	"bytes"
	"io"
	"log/slog" // Use the new structured log package
	"net/http"
	"os"
	infrastructure_logger "passport-mrs-go-sign-up-service/infrastructure/logger"
	"strings"
	"time"
)

func RequestLogger(next http.Handler) http.Handler {
	// Initialize slog to write to stdout
	// You can change LevelDebug to LevelInfo to hide debug logs
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var bodyLog string = "<non-json or empty>"

		// 1. Generate and attach Trace ID to context
		ctx := infrastructure_logger.WithTraceID(r.Context())
		traceID := infrastructure_logger.FromContext(ctx)

		// 2. Add Trace ID to the Response Header so the client can see it too
		w.Header().Set("X-Trace-ID", traceID)

		// 3. Update the request with the new context
		r = r.WithContext(ctx)

		contentType := r.Header.Get("Content-Type")
		if strings.HasPrefix(contentType, "application/json") {
			bodyBytes, err := io.ReadAll(r.Body)
			if err == nil && len(bodyBytes) > 0 {
				bodyLog = string(bodyBytes)
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		start := time.Now()
		next.ServeHTTP(w, r)

		// Using Structured Logging with Levels
		logger.Info("request processed",
			"method", r.Method,
			"url", r.URL.Path,
			"duration", time.Since(start),
			"trace_id", traceID,
		)

		// This will only show up if opts.Level is set to slog.LevelDebug
		logger.Debug("request details",
			"body", bodyLog,
			"content_type", contentType,
		)
	})
}
