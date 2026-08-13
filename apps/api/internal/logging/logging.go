package logging

import (
	"log/slog"
	"os"
)

// Setup configures JSON structured logging to stdout.
// Full LOG_LEVEL / request correlation support is tracked in E12-F5-T1
// (see docs/adr/0002-api-observability.md).
func Setup() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}
