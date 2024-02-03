package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"renting/internal/config"
	"renting/internal/repository"
	"renting/internal/repository/postgres"
	"renting/internal/server"
	"renting/models"
)

type CustomLogger struct {
	*slog.Logger
}

func NewCustomLogger() *CustomLogger {
	return &CustomLogger{
		slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		),
	}
}

func (l *CustomLogger) Info(message string) {
	l.Logger.Log(context.Background(), slog.LevelInfo, message)
}

func (l *CustomLogger) Debug(message string) {
	l.Logger.Log(context.Background(), slog.LevelDebug, message)
}

func (l *CustomLogger) Error(message string) {
	l.Logger.Log(context.Background(), slog.LevelError, message)
}

func main() {
	log := NewCustomLogger()
	log.Info("Starting the application")

	store := postgres.NewPostgresDb(postgres.GetDBInstance())
	cfg := config.NewConfig(store, log.Logger)

	app := server.NewServer(cfg, ":3000")
	log.Info("Starting server on port 3000")

	app.Run()

	//просто пример
	username := "_user"
	password := "_password"
	authenticated := authenticateUser(store, log, username, password)

	if authenticated {
		log.Info(fmt.Sprintf("User '%s' has been successfully authenticated", username))
	} else {
		log.Error(fmt.Sprintf("Failed to authenticate user '%s'", username))
	}
}

func authenticateUser(store repository.Store, log *CustomLogger, username, password string) bool {
	user, err := store.LoginUser(username, password)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to authenticate user: %v", err))
		return false
	}

	return user != (models.User{})
}
