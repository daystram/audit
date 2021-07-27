package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/daystram/audit/audit-be/constants"
	"github.com/daystram/audit/audit-be/datatransfers"
	"github.com/daystram/audit/audit-be/handlers"
)

func GETServiceList(h handlers.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var err error
		var serviceInfos []datatransfers.ServiceInfo
		if serviceInfos, err = h.ServiceGetAll(c.Param("application_id")); err != nil {
			c.JSON(http.StatusInternalServerError, datatransfers.Response{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, datatransfers.Response{
			Data: serviceInfos,
		})
	})
}

func POSTServiceCreate(h handlers.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var err error
		var serviceInfo datatransfers.ServiceInfo
		if err = c.ShouldBindJSON(&serviceInfo); err != nil {
			c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
			return
		}
		serviceInfo.ApplicationID = c.Param("application_id")
		if serviceInfo.Type != constants.ServiceTypeHTTP && serviceInfo.Type != constants.ServiceTypeTCP && serviceInfo.Type != constants.ServiceTypePING {
			c.JSON(http.StatusBadRequest, datatransfers.Response{
				Error: fmt.Sprintf("unknown service type %s", serviceInfo.Type),
			})
			return
		}
		// TODO: validate Config
		if _, err = h.ServiceCreate(serviceInfo); err != nil {
			c.JSON(http.StatusInternalServerError, datatransfers.Response{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, datatransfers.Response{})
	})
}

func GETServiceDetail(h handlers.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var err error
		var serviceInfo datatransfers.ServiceInfo
		if serviceInfo, err = h.ServiceGetOne(c.Param("service_id"), c.Param("application_id")); err != nil {
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
			Data: serviceInfo,
		})
	})
}

func PUTServiceUpdate(h handlers.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var err error
		var serviceInfo datatransfers.ServiceInfo
		if err = c.ShouldBindJSON(&serviceInfo); err != nil {
			c.JSON(http.StatusBadRequest, datatransfers.Response{Error: err.Error()})
			return
		}
		serviceInfo.ID = c.Param("service_id")
		serviceInfo.ApplicationID = c.Param("application_id")
		if serviceInfo.Type != constants.ServiceTypeHTTP && serviceInfo.Type != constants.ServiceTypeTCP && serviceInfo.Type != constants.ServiceTypePING {
			c.JSON(http.StatusBadRequest, datatransfers.Response{
				Error: fmt.Sprintf("unknown service type %s", serviceInfo.Type),
			})
			return
		}
		// TODO: validate Config
		if err = h.ServiceUpdate(serviceInfo); err != nil {
			c.JSON(http.StatusInternalServerError, datatransfers.Response{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, datatransfers.Response{})
	})
}

func DELETEServiceDelete(h handlers.HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		var err error
		if err = h.ServiceDelete(c.Param("service_id"), c.Param("application_id")); err != nil {
			c.JSON(http.StatusInternalServerError, datatransfers.Response{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, datatransfers.Response{})
	})
}
