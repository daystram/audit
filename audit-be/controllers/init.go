package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/daystram/audit/audit-be/handlers"
)

func InitializeRouter(h handlers.HandlerFunc) (router *gin.Engine) {
	router = gin.Default()
	api := router.Group("/api")
	{
		monitor := api.Group("/monitor")
		{
			monitor.GET("/", GETMonitorList(h))
		}

		application := api.Group("/application")
		{
			application.GET("/", GETApplicationList(h))
			application.POST("/", POSTApplicationCreate(h))
			application.GET("/:application_id", GETApplicationDetail(h))
			application.PUT("/:application_id", PUTApplicationUpdate(h))
			application.DELETE("/:application_id", DELETEApplicationDelete(h))

			service := application.Group("/:application_id/service")
			{
				service.GET("/", GETServiceList(h))
				service.POST("/", POSTServiceCreate(h))
				service.GET("/:service_id", GETServiceDetail(h))
				service.PUT("/:service_id", PUTServiceUpdate(h))
				service.DELETE("/:service_id", DELETEServiceDelete(h))
			}
		}
	}
	return
}
