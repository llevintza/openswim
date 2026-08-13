package config

import "testing"

func TestLoadDefaultsPort(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("DATABASE_URL", "postgres://u:p@localhost:5432/openswim?sslmode=disable")
	t.Setenv("JWT_SECRET", "test-secret")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Port != "8080" {
		t.Fatalf("Port = %q, want 8080", cfg.Port)
	}
}

func TestLoadRequiresSecrets(t *testing.T) {
	t.Setenv("PORT", "9090")
	t.Setenv("DATABASE_URL", "")
	t.Setenv("JWT_SECRET", "")

	_, err := Load()
	if err == nil {
		t.Fatal("expected error for missing env")
	}
}

func TestLoadOK(t *testing.T) {
	t.Setenv("PORT", "9090")
	t.Setenv("DATABASE_URL", "postgres://u:p@localhost:5432/openswim?sslmode=disable")
	t.Setenv("JWT_SECRET", "test-secret")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Port != "9090" {
		t.Fatalf("Port = %q, want 9090", cfg.Port)
	}
	if cfg.DatabaseURL == "" || cfg.JWTSecret == "" {
		t.Fatal("expected database URL and JWT secret")
	}
}
