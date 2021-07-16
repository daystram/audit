package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/daystram/audit/audit-be/datatransfers"
	"github.com/daystram/audit/audit-be/handlers"
)

func GETApplicationList(h handlers.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var err error
		var applicationInfos []datatransfers.ApplicationInfo
		if applicationInfos, err = h.ApplicationGetAll(); err != nil {
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

func POSTApplicationCreate(h handlers.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var err error
		var applicationInfo datatransfers.ApplicationInfo
		if err = c.ShouldBindJSON(&applicationInfo); err != nil {
			c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
			return
		}
		if _, err = handlers.Handler.ApplicationCreate(applicationInfo); err != nil {
			c.JSON(http.StatusInternalServerError, datatransfers.Response{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, datatransfers.Response{})
	})
}

func GETApplicationDetail(h handlers.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var err error
		var applicationInfo datatransfers.ApplicationInfo
		if applicationInfo, err = handlers.Handler.ApplicationGetOne(c.Param("application_id")); err != nil {
			c.JSON(http.StatusInternalServerError, datatransfers.Response{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, datatransfers.Response{
			Data: applicationInfo,
		})
	})
}

func PUTApplicationUpdate(h handlers.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var err error
		var applicationInfo datatransfers.ApplicationInfo
		if err = c.ShouldBindJSON(&applicationInfo); err != nil {
			c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
			return
		}
		if err = handlers.Handler.ApplicationUpdate(applicationInfo); err != nil {
			c.JSON(http.StatusInternalServerError, datatransfers.Response{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, datatransfers.Response{})
	})
}

func DELETEApplicationDelete(h handlers.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var err error
		if err = handlers.Handler.ApplicationDelete(c.Param("application_id")); err != nil {
			c.JSON(http.StatusInternalServerError, datatransfers.Response{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, datatransfers.Response{})
	})
}
