package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/daystram/audit/audit-be/handlers"
)

func InitializeRouter(h handlers.HandlerFunc) (router *gin.Engine) {
	router = gin.Default()
	application := router.Group("/application")
	{
		application.GET("/", GETApplicationList(h))
		application.POST("/", POSTApplicationCreate(h))
		application.GET("/:application_id", GETApplicationDetail(h))
		application.PUT("/:application_id", PUTApplicationUpdate(h))
		application.DELETE("/:application_id", DELETEApplicationDelete(h))
	}
	return
}
