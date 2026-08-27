package config

import "testing"

func TestLoadRequiresDatabaseURL(t *testing.T) {
	t.Setenv("DATABASE_URL", "")
	t.Setenv("ADMIN_API_TOKEN", "a-valid-local-admin-token-with-32-chars")

	_, err := Load()
	if err == nil {
		t.Fatal("Load() error = nil, want error when DATABASE_URL is absent")
	}
}

func TestLoadUsesDefaults(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://user:password@localhost:5432/pera_api")
	t.Setenv("ADMIN_API_TOKEN", "a-valid-local-admin-token-with-32-chars")
	t.Setenv("HTTP_ADDR", "")
	t.Setenv("DATABASE_MAX_CONNS", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.HTTPAddr != defaultHTTPAddr {
		t.Fatalf("HTTPAddr = %q, want %q", cfg.HTTPAddr, defaultHTTPAddr)
	}
	if cfg.DatabaseMaxConns != defaultDatabaseMaxConns {
		t.Fatalf("DatabaseMaxConns = %d, want %d", cfg.DatabaseMaxConns, defaultDatabaseMaxConns)
	}
}

func TestLoadRejectsInvalidConnectionLimit(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://user:password@localhost:5432/pera_api")
	t.Setenv("ADMIN_API_TOKEN", "a-valid-local-admin-token-with-32-chars")
	t.Setenv("DATABASE_MAX_CONNS", "0")

	_, err := Load()
	if err == nil {
		t.Fatal("Load() error = nil, want error for invalid DATABASE_MAX_CONNS")
	}
}
