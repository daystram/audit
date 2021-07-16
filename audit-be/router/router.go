package router

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/daystram/audit/audit-be/controllers/v1"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	apiV1 := router.Group("/api/v1")
	{
		application := apiV1.Group("/application")
		{
			application.GET("/", v1.GETApplicationList)
			application.POST("/", v1.POSTApplicationCreate)
			application.GET("/:application_id", v1.GETApplicationDetail)
			application.PUT("/:application_id", v1.PUTApplicationUpdate)
			application.DELETE("/:application_id", v1.DELETEApplicationDelete)
		}
	}
	return
}
