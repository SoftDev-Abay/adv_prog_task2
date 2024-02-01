package config

import (
	"log/slog"
	"renting/internal/repository"
)

type Config struct {
	Store repository.Store
	Log   *slog.Logger
}

func NewConfig(store repository.Store, logger *slog.Logger) *Config {
	return &Config{
		Store: store,
		Log:   logger,
	}
}
