package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

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
		if _, err = h.ApplicationCreate(applicationInfo); err != nil {
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
		if applicationInfo, err = h.ApplicationGetOne(c.Param("application_id")); err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, datatransfers.Response{})
			} else {
				c.JSON(http.StatusInternalServerError, datatransfers.Response{
					Error: err.Error(),
				})
			}
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
		applicationInfo.ID = c.Param("application_id")
		if err = h.ApplicationUpdate(applicationInfo); err != nil {
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
		if err = h.ApplicationDelete(c.Param("application_id")); err != nil {
			c.JSON(http.StatusInternalServerError, datatransfers.Response{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, datatransfers.Response{})
	})
}
