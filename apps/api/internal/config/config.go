package config

import (
	"fmt"
	"os"
	"strings"
)

// Config holds process configuration loaded from the environment.
type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
}

// Load reads required settings from the environment.
// PORT defaults to 8080; DATABASE_URL and JWT_SECRET are required.
func Load() (Config, error) {
	cfg := Config{
		Port:        strings.TrimSpace(os.Getenv("PORT")),
		DatabaseURL: strings.TrimSpace(os.Getenv("DATABASE_URL")),
		JWTSecret:   strings.TrimSpace(os.Getenv("JWT_SECRET")),
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	var missing []string
	if cfg.DatabaseURL == "" {
		missing = append(missing, "DATABASE_URL")
	}
	if cfg.JWTSecret == "" {
		missing = append(missing, "JWT_SECRET")
	}
	if len(missing) > 0 {
		return Config{}, fmt.Errorf("missing required env: %s", strings.Join(missing, ", "))
	}
	return cfg, nil
}
