package router

import (
	"github.com/gin-gonic/gin"

	"github.com/daystram/go-gin-gorm-boilerplate/controllers/middleware"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	v1route := router.Group("/api/v1")
	v1route.Use(
		middleware.CORSMiddleware,
		middleware.AuthMiddleware,
	)
	return
}
