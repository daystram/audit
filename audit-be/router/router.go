package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	return
}
