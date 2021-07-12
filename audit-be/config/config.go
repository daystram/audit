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

	DBHost     string
	DBPort     int
	DBDatabase string
	DBUsername string
	DBPassword string

	InfluxDBURL          string
	InfluxDBToken        string
	InfluxDBOrganization string
	InfluxDBBucket       string
}

func InitializeAppConfig() {
	viper.SetConfigName(".env") // for development
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

	AppConfig.DBHost = viper.GetString("DB_HOST")
	AppConfig.DBPort = viper.GetInt("DB_PORT")
	AppConfig.DBDatabase = viper.GetString("DB_DATABASE")
	AppConfig.DBUsername = viper.GetString("DB_USERNAME")
	AppConfig.DBPassword = viper.GetString("DB_PASSWORD")

	AppConfig.InfluxDBURL = viper.GetString("INFLUXDB_URL")
	AppConfig.InfluxDBToken = viper.GetString("INFLUXDB_TOKEN")
	AppConfig.InfluxDBOrganization = viper.GetString("INFLUXDB_ORGANIZATION")
	AppConfig.InfluxDBBucket = viper.GetString("INFLUXDB_BUCKET")

	log.Printf("[INIT] Configuration loaded!")
}
