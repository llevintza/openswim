package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/llevintza/openswim/apps/api/internal/config"
	"github.com/llevintza/openswim/apps/api/internal/db"
	"github.com/llevintza/openswim/apps/api/internal/httpserver"
	"github.com/llevintza/openswim/apps/api/internal/logging"
	"github.com/llevintza/openswim/apps/api/internal/migrate"
)

func main() {
	logger := logging.Setup()
	if err := run(logger); err != nil {
		logger.Error("api exited", "error", err)
		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	pool, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer pool.Close()

	migrationsDir, err := migrationsPath()
	if err != nil {
		return err
	}
	if err := migrate.Up(cfg.DatabaseURL, migrationsDir); err != nil {
		return err
	}

	srv := httpserver.New(pool)
	httpServer := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           srv.Handler(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info("listening", "addr", httpServer.Addr)
		errCh <- httpServer.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("shutdown: %w", err)
		}
		return nil
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	}
}

func migrationsPath() (string, error) {
	// Prefer migrations next to the module root when running from apps/api.
	candidates := []string{
		"migrations",
		filepath.Join("apps", "api", "migrations"),
	}
	for _, c := range candidates {
		if info, err := os.Stat(c); err == nil && info.IsDir() {
			abs, err := filepath.Abs(c)
			if err != nil {
				return "", err
			}
			return abs, nil
		}
	}
	return "", fmt.Errorf("migrations directory not found")
}
