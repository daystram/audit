package router

import (
	"github.com/gin-gonic/gin"

	"github.com/daystram/go-gin-gorm-boilerplate/controllers/middleware"
	v1 "github.com/daystram/go-gin-gorm-boilerplate/controllers/v1"
	"github.com/daystram/go-gin-gorm-boilerplate/utils"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	v1route := router.Group("/api/v1")
	v1route.Use(
		middleware.CORSMiddleware,
		middleware.AuthMiddleware,
	)
	{
		user := v1route.Group("/user")
		{
			user.GET("/:username", utils.AuthOnly, v1.GETUser)
			user.PUT("", utils.AuthOnly, v1.PUTUser)
		}
	}
	return
}
