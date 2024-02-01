package main

import (
	"log/slog"
	"os"
	"renting/internal/config"
	"renting/internal/repository/postgres"
	"renting/internal/server"
)

const port = 3000

func main() {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	store := postgres.NewPostgresDb(postgres.GetDBInstance())

	cfg := config.NewConfig(store, log)

	app := server.NewServer(cfg, ":3000")
	app.Run()
}
