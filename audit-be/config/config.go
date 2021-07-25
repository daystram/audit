package config

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	Port        int
	Environment string
	Debug       bool

	TrackerServerPort int

	RatifyIssuer       string
	RatifyClientID     string
	RatifyClientSecret string

	DBHost     string
	DBPort     int
	DBDatabase string
	DBUsername string
	DBPassword string

	InfluxDBURL             string
	InfluxDBToken           string
	InfluxDBOrganization    string
	InfluxDBReportsBucket   string
	InfluxDBIncidentsBucket string
}

func InitializeAppConfig() {
	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	AppConfig.Port = viper.GetInt("PORT")
	AppConfig.Environment = viper.GetString("ENVIRONMENT")
	AppConfig.Debug = viper.GetBool("DEBUG")

	AppConfig.TrackerServerPort = viper.GetInt("TRACKER_SERVER_PORT")

	AppConfig.RatifyIssuer = viper.GetString("RATIFY_ISSUER")
	AppConfig.RatifyClientID = viper.GetString("RATIFY_CLIENT_ID")
	AppConfig.RatifyClientSecret = viper.GetString("RATIFY_CLIENT_SECRET")

	AppConfig.DBHost = viper.GetString("DB_HOST")
	AppConfig.DBPort = viper.GetInt("DB_PORT")
	AppConfig.DBDatabase = viper.GetString("DB_DATABASE")
	AppConfig.DBUsername = viper.GetString("DB_USERNAME")
	AppConfig.DBPassword = viper.GetString("DB_PASSWORD")

	AppConfig.InfluxDBURL = viper.GetString("INFLUXDB_URL")
	AppConfig.InfluxDBToken = viper.GetString("INFLUXDB_TOKEN")
	AppConfig.InfluxDBOrganization = viper.GetString("INFLUXDB_ORGANIZATION")
	AppConfig.InfluxDBReportsBucket = viper.GetString("INFLUXDB_REPORTS_BUCKET")
	AppConfig.InfluxDBIncidentsBucket = viper.GetString("INFLUXDB_INCIDENTS_BUCKET")

	log.Printf("[INIT] Configuration loaded!")
}
