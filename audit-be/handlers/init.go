package handlers

import (
	"context"
	"fmt"
	"log"

	influxlib "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/daystram/go-gin-gorm-boilerplate/config"
)

var Handler HandlerFunc

type HandlerFunc interface{}

type module struct {
	db     *dbEntity
	influx *influxEntity
}

type dbEntity struct {
	conn *gorm.DB
}

type influxEntity struct {
	conn     *influxlib.Client
	writeAPI api.WriteAPI
	queryAPI api.QueryAPI
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
		log.Fatalf("[INIT] Failed connecting to PostgreSQL at %s:%d. %+v\n", config.AppConfig.DBHost, config.AppConfig.DBPort, err)
	}
	log.Printf("[INIT] Successfully connected to PostgreSQL\n")

	// Initialize InfluxDB
	var ready bool
	influx := influxlib.NewClient(config.AppConfig.InfluxDBURL, config.AppConfig.InfluxDBToken)
	ready, err = influx.Ready(context.Background())
	if err != nil {
		log.Fatalf("[INIT] Failed connecting to InfluxDB at %s. %+v\n", config.AppConfig.InfluxDBURL, err)
	}
	if !ready {
		log.Fatalf("[INIT] Failed connecting to InfluxDB at %s. influxdb instance not ready\n", config.AppConfig.InfluxDBURL)
	}
	log.Printf("[INIT] Successfully connected to InfluxDB\n")

	// Compose handler modules
	Handler = &module{
		db: &dbEntity{
			conn: db,
		},
		influx: &influxEntity{
			conn:     &influx,
			writeAPI: influx.WriteAPI(config.AppConfig.InfluxDBOrganization, config.AppConfig.InfluxDBBucket),
			queryAPI: influx.QueryAPI(config.AppConfig.InfluxDBOrganization),
		},
	}
}
