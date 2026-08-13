package migrate

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Up applies all pending SQL migrations from migrationsDir.
// With no *.sql files this is a no-op stub until the first schema lands.
func Up(databaseURL, migrationsDir string) error {
	hasSQL, err := hasSQLMigrations(migrationsDir)
	if err != nil {
		return err
	}
	if !hasSQL {
		return nil
	}

	sourceURL := "file://" + migrationsDir
	m, err := migrate.New(sourceURL, withPGXScheme(databaseURL))
	if err != nil {
		return fmt.Errorf("migrate open: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrate up: %w", err)
	}
	return nil
}

func hasSQLMigrations(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, fmt.Errorf("read migrations dir: %w", err)
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if strings.HasSuffix(e.Name(), ".sql") {
			return true, nil
		}
	}
	return false, nil
}

func withPGXScheme(databaseURL string) string {
	switch {
	case strings.HasPrefix(databaseURL, "postgres://"):
		return "pgx5://" + strings.TrimPrefix(databaseURL, "postgres://")
	case strings.HasPrefix(databaseURL, "postgresql://"):
		return "pgx5://" + strings.TrimPrefix(databaseURL, "postgresql://")
	case strings.HasPrefix(databaseURL, "pgx5://"):
		return databaseURL
	default:
		return databaseURL
	}
}
