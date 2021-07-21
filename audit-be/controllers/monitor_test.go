package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/daystram/audit/audit-be/datatransfers"
	mock_handlers "github.com/daystram/audit/audit-be/mocks/handlers"
)

type V1MonitorTestSuite struct {
	suite.Suite
	Router      *gin.Engine
	MockHandler *mock_handlers.MockHandlerFunc
}

func (suite *V1MonitorTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.MockHandler = mock_handlers.NewMockHandlerFunc(ctrl)
	suite.Router = InitializeRouter(suite.MockHandler)
	gin.SetMode(gin.TestMode)
}

func (suite *V1MonitorTestSuite) TestGETMonitorList() {
	suite.T().Run("applications exist", func(t *testing.T) {
		suite.MockHandler.EXPECT().MonitorGetAll().Return([]datatransfers.ApplicationInfo{{
			ID: "app_id",
			Services: []datatransfers.ServiceInfo{{
				ID:            "service_id",
				ApplicationID: "app_id",
				Name:          "Test Service",
				Description:   "Description",
				Endpoint:      "https://test.daystram.com",
				Type:          "http",
				Config:        "{}",
				Enabled:       true,
				Showcase:      true,
				CreatedAt:     1,
				UpdatedAt:     2,
			}},
			Name:        "Test App",
			Description: "Description",
			CreatedAt:   1,
			UpdatedAt:   2,
		}}, nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/monitor/", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.JSONEq(suite.T(), `{
			"data": [{
				"id": "app_id",
				"services": [{
					"id": "service_id",
					"name": "Test Service",
					"description": "Description",
					"endpoint": "https://test.daystram.com",
					"type": "http",
					"config": "{}",
					"enabled": true,
					"showcase": true,
					"createdAt": 1,
					"updatedAt": 2
				}],
				"name": "Test App",
				"description": "Description",
				"createdAt": 1,
				"updatedAt": 2
			}]
		}`, w.Body.String())
	})
	suite.T().Run("no applications", func(t *testing.T) {
		suite.MockHandler.EXPECT().MonitorGetAll().Return(make([]datatransfers.ApplicationInfo, 0), nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/monitor/", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.JSONEq(suite.T(), `{"data": []}`, w.Body.String())
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockHandler.EXPECT().MonitorGetAll().Return([]datatransfers.ApplicationInfo{}, errors.New(""))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/monitor/", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Contains(suite.T(), w.Body.String(), `"error"`)
	})
}

func TestV1MonitorTestSuite(t *testing.T) {
	suite.Run(t, new(V1MonitorTestSuite))
}
