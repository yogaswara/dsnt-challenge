package logger

import (
	"log/slog"
	"os"
)

// Init initializes the default standard logger with JSON format
func Init() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

// Info logs an informational message
func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

// Error logs an error message
func Error(msg string, err error, args ...any) {
	if err != nil {
		args = append(args, "error", err.Error())
	}
	slog.Error(msg, args...)
}
