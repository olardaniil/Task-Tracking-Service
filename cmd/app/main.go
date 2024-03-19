package main

import (
	"log"
	"task_tracking_service/configs"
	"task_tracking_service/internal/handler"
	"task_tracking_service/internal/repository"
	"task_tracking_service/internal/service"
)

// @title Task-Tracking-Service
// @version 1.0
// @description API Server for FilmLibraryService

// @host 127.0.0.1:8080
// @BasePath /api

func main() {
	db, err := repository.NewPostgresDB(configs.NewConfig())
	if err != nil {
		log.Fatalf("Ошибка при инициализации БД: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes(configs.NewConfig().AppPort)
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
}
