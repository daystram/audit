package handlers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/daystram/go-gin-gorm-boilerplate/config"
)

var Handler HandlerFunc

type HandlerFunc interface{}

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn *gorm.DB
}

func InitializeHandler() {
	var err error

	// Initialize DB
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
			config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBDatabase,
			config.AppConfig.DBUsername, config.AppConfig.DBPassword),
	), &gorm.Config{})
	if err != nil {
		log.Fatalf("[INIT] Failed connecting to PostgreSQL Database at %s:%d. %+v\n",
			config.AppConfig.DBHost, config.AppConfig.DBPort, err)
	}
	log.Printf("[INIT] Successfully connected to PostgreSQL Database\n")

	// Compose handler modules
	Handler = &module{
		db: &dbEntity{
			conn: db,
		},
	}
}
