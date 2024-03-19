package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"task_tracking_service/configs"
)

func NewPostgresDB(cfg configs.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPass, cfg.DBSSLMode))
	if err != nil {
		return nil, err
	}
	//
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	// Создание пула соединений
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	//
	log.Println("Подключение к БД прошло успешно")
	return db, nil
}
