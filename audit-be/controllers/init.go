package controllers

import (
	"github.com/gin-gonic/gin"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	application := router.Group("/application")
	{
		application.GET("/", GETApplicationList)
		application.POST("/", POSTApplicationCreate)
		application.GET("/:application_id", GETApplicationDetail)
		application.PUT("/:application_id", PUTApplicationUpdate)
		application.DELETE("/:application_id", DELETEApplicationDelete)
	}
	return
}
