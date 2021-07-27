package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/daystram/audit/audit-be/config"
	"github.com/daystram/audit/audit-be/controllers"
	"github.com/daystram/audit/audit-be/handlers"
)

func init() {
	config.InitializeAppConfig()
	if !config.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	h, err := handlers.InitializeHandler()
	if err != nil {
		log.Fatal(err)
	}
	if err = h.InitializeTrackerServer(config.AppConfig.TrackerServerPort); err != nil {
		log.Fatal(err)
	}
	if err = h.InitializeScheduler(); err != nil {
		log.Fatal(err)
	}

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        controllers.InitializeRouter(h),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
