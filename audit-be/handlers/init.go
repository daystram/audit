package handlers

import (
	"context"
	"fmt"
	"log"

	influxlib "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/daystram/audit/audit-be/config"
	"github.com/daystram/audit/audit-be/datatransfers"
	"github.com/daystram/audit/audit-be/models"
)

type HandlerFunc interface {
	// Application
	ApplicationGetAll() (applicationInfos []datatransfers.ApplicationInfo, err error)
	ApplicationGetOne(applicationID string) (applicationInfo datatransfers.ApplicationInfo, err error)
	ApplicationCreate(applicationInfo datatransfers.ApplicationInfo) (applicationID string, err error)
	ApplicationUpdate(applicationInfo datatransfers.ApplicationInfo) (err error)
	ApplicationDelete(applicationID string) (err error)

	// Service
	ServiceGetAll(applicationID string) (serviceInfos []datatransfers.ServiceInfo, err error)
	ServiceGetOne(serviceID, applicationID string) (applicationInfo datatransfers.ServiceInfo, err error)
	ServiceCreate(serviceInfo datatransfers.ServiceInfo) (serviceID string, err error)
	ServiceUpdate(serviceInfo datatransfers.ServiceInfo) (err error)
	ServiceDelete(serviceID, applicationID string) (err error)

	// Tracker
	InitializeTrackerServer(port int)
}

type module struct {
	db            *dbEntity
	influx        *influxEntity
	trackerServer TrackerServer
}

type dbEntity struct {
	conn             *gorm.DB
	applicationOrmer models.ApplicationOrmer
	serviceOrmer     models.ServiceOrmer
}

type influxEntity struct {
	conn     *influxlib.Client
	writeAPI api.WriteAPI
	queryAPI api.QueryAPI
}

func InitializeHandler() (handler *module, err error) {
	// Initialize DB
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
			config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBDatabase,
			config.AppConfig.DBUsername, config.AppConfig.DBPassword),
	), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed connecting to PostgreSQL at %s:%d. %+v", config.AppConfig.DBHost, config.AppConfig.DBPort, err)
	}
	log.Printf("[INIT] Successfully connected to PostgreSQL\n")

	// Initialize InfluxDB
	var ready bool
	influx := influxlib.NewClient(config.AppConfig.InfluxDBURL, config.AppConfig.InfluxDBToken)
	ready, err = influx.Ready(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed connecting to InfluxDB at %s. %+v", config.AppConfig.InfluxDBURL, err)
	}
	if !ready {
		return nil, fmt.Errorf("failed connecting to InfluxDB at %s. influxdb instance not ready", config.AppConfig.InfluxDBURL)
	}
	log.Printf("[INIT] Successfully connected to InfluxDB\n")

	// Compose handler modules
	handler = &module{
		db: &dbEntity{
			conn:             db,
			applicationOrmer: models.NewApplicationOrmer(db),
			serviceOrmer:     models.NewServiceOrmer(db),
		},
		influx: &influxEntity{
			conn:     &influx,
			writeAPI: influx.WriteAPI(config.AppConfig.InfluxDBOrganization, config.AppConfig.InfluxDBBucket),
			queryAPI: influx.QueryAPI(config.AppConfig.InfluxDBOrganization),
		},
	}
	return
}
