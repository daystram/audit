package config

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	TrackerServer string
	TrackerID     string
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

	AppConfig.TrackerServer = viper.GetString("TRACKER_SERVER")
	AppConfig.TrackerID = viper.GetString("TRACKER_ID")

	log.Printf("[INIT] Configuration loaded!")
}
