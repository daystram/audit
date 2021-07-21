package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/daystram/audit/audit-be/datatransfers"
	"github.com/daystram/audit/audit-be/handlers"
)

func GETMonitorList(h handlers.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var err error
		var applicationInfos []datatransfers.ApplicationInfo
		if applicationInfos, err = h.MonitorGetAll(); err != nil {
			c.JSON(http.StatusInternalServerError, datatransfers.Response{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, datatransfers.Response{
			Data: applicationInfos,
		})
	})
}
