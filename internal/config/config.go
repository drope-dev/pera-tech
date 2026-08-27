package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	defaultHTTPAddr         = ":8080"
	defaultDatabaseMaxConns = int32(4)
)

type Config struct {
	HTTPAddr         string
	DatabaseURL      string
	DatabaseMaxConns int32
	AdminAPIToken    string
}

func Load() (Config, error) {
	databaseURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok || strings.TrimSpace(databaseURL) == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}
	adminAPIToken, ok := os.LookupEnv("ADMIN_API_TOKEN")
	if !ok || len(strings.TrimSpace(adminAPIToken)) < 32 {
		return Config{}, fmt.Errorf("ADMIN_API_TOKEN must contain at least 32 characters")
	}

	maxConns, err := int32FromEnv("DATABASE_MAX_CONNS", defaultDatabaseMaxConns)
	if err != nil {
		return Config{}, err
	}

	return Config{
		HTTPAddr:         valueOrDefault("HTTP_ADDR", defaultHTTPAddr),
		DatabaseURL:      databaseURL,
		DatabaseMaxConns: maxConns,
		AdminAPIToken:    adminAPIToken,
	}, nil
}

func valueOrDefault(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func int32FromEnv(key string, fallback int32) (int32, error) {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback, nil
	}

	parsed, err := strconv.ParseInt(value, 10, 32)
	if err != nil || parsed < 1 || parsed > 50 {
		return 0, fmt.Errorf("%s must be an integer between 1 and 50", key)
	}
	return int32(parsed), nil
}
