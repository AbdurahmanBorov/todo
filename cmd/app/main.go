package main

import (
	"context"
	"log"
	"todo/internal/api"
	"todo/internal/config"
	db2 "todo/internal/pkg/db"
	"todo/internal/repository"
	service2 "todo/internal/service"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatalf("Failed load configuration: %v", err)
	}

	db, err := db2.OpenDB(context.Background(), cfg.DBConfig)
	if err != nil {
		log.Fatalf("Failed connect DB: %v", err)
	}

	err = db2.Migrate(db, cfg.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(db)
	service := service2.New(repo)
	router := api.New(service)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
