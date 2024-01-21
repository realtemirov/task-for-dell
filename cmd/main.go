package main

import (
	"fmt"
	"log"

	"github.com/realtemirov/task-for-dell/config"
	"github.com/realtemirov/task-for-dell/internal/server"
	"github.com/realtemirov/task-for-dell/pkg/db/postgres"
	"github.com/realtemirov/task-for-dell/pkg/logger"
)

// @title Blog and News API.
// @version 1.0
// @description Blog and News API Server.
// @contact.name Jakhongir Temirov
// @contact.url https://github.com/realtemirov
// @contact.email realjakhongir@gmail.com
// @BasePath /v1
func main() {

	fmt.Println("Starting application...")
	cfg, err := config.LoadConfig("./config/config-local")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	fmt.Println(cfg)
	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()

	logger.Info("Application started")
	logger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		logger.Fatalf("failed to connect to db: %v", err)
	} else {
		logger.Info("successfully connected to db")
	}
	defer db.Close()

	serv := server.NewServer(cfg, logger, db)
	if err = serv.Run(); err != nil {
		log.Fatal(err)
	}
}
