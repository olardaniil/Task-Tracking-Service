package main

import (
	"github.com/joho/godotenv"
	"log"
	"task_tracking_service/configs"
	"task_tracking_service/internal/handler"
	"task_tracking_service/internal/repository"
	"task_tracking_service/internal/service"
)

//	@title		Task-Tracking-Service
//	@version	1.0

//	@BasePath	/api

func main() {
	cfg, err := configs.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Ошибка при инициализации БД: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes(cfg.AppPort)
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
}
