package main

import (
	"log"

	"google.golang.org/grpc"

	"github.com/daystram/audit/audit-tr/config"
	"github.com/daystram/audit/audit-tr/handler"
)

func main() {
	config.InitializeAppConfig()

	conn, err := grpc.Dial(config.AppConfig.TrackerServer, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[INIT] failed dialling audit-be at %s. %v", config.AppConfig.TrackerServer, err)
	}
	defer conn.Close()

	h, err := handler.InitializeHandler(conn)
	if err != nil {
		log.Fatal(err)
	}
	h.SubscribeTracking()
}
