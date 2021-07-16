package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/daystram/audit/audit-be/datatransfers"
	mock_handlers "github.com/daystram/audit/audit-be/mocks/handlers"
)

type V1ServiceTestSuite struct {
	suite.Suite
	Router      *gin.Engine
	MockHandler *mock_handlers.MockHandlerFunc
}

func (suite *V1ServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.MockHandler = mock_handlers.NewMockHandlerFunc(ctrl)
	suite.Router = InitializeRouter(suite.MockHandler)
	gin.SetMode(gin.ReleaseMode)
}

func (suite *V1ServiceTestSuite) TestGETServiceList() {
	suite.T().Run("applications exist", func(t *testing.T) {
		suite.MockHandler.EXPECT().ServiceGetAll("app_id").Return([]datatransfers.ServiceInfo{{
			ID:            "service_id",
			ApplicationID: "app_id",
			Name:          "Test Service",
			Description:   "Description",
			Endpoint:      "https://service.daystram.com",
			Type:          "http",
			Config:        "{}",
			Showcase:      true,
			CreatedAt:     1,
			UpdatedAt:     2,
		}}, nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/application/app_id/service/", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.JSONEq(suite.T(), `{"data":[{"id":"service_id", "name":"Test Service", "description": "Description", 
		"endpoint":"https://service.daystram.com", "type":"http", "config":"{}", "showcase":true, "createdAt": 1, "updatedAt": 2}]}`, w.Body.String())
	})
	suite.T().Run("no services", func(t *testing.T) {
		suite.MockHandler.EXPECT().ServiceGetAll("app_id").Return(make([]datatransfers.ServiceInfo, 0), nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/application/app_id/service/", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.JSONEq(suite.T(), `{"data": []}`, w.Body.String())
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockHandler.EXPECT().ServiceGetAll("app_id").Return([]datatransfers.ServiceInfo{}, errors.New(""))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/application/app_id/service/", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Contains(suite.T(), w.Body.String(), `"error"`)
	})
}

func (suite *V1ServiceTestSuite) TestPOSTServiceCreate() {
	suite.T().Run("create", func(t *testing.T) {
		suite.MockHandler.EXPECT().ServiceCreate(datatransfers.ServiceInfo{
			ApplicationID: "app_id",
			Name:          "Test Service",
			Description:   "Description",
			Endpoint:      "https://service.daystram.com",
			Type:          "http",
			Config:        "{}",
			Showcase:      true,
		}).Return("service_id", nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/application/app_id/service/", strings.NewReader(`{"name":"Test Service", "description": "Description", 
		"endpoint":"https://service.daystram.com", "type":"http", "config":"{}", "showcase":true}`))
		req.Header.Set("Content-Type", "application/json")
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.JSONEq(suite.T(), `{}`, w.Body.String())
	})
	suite.T().Run("bad request", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/application/app_id/service/", strings.NewReader(`{"description": "Description"}`))
		req.Header.Set("Content-Type", "application/json")
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	})
	suite.T().Run("invalid type", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/application/app_id/service/", strings.NewReader(`{"name":"Test Service", "description": "Description", 
		"endpoint":"https://service.daystram.com", "type":"whattypeisthis", "config":"{}", "showcase":true}`))
		req.Header.Set("Content-Type", "application/json")
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockHandler.EXPECT().ServiceCreate(datatransfers.ServiceInfo{
			ApplicationID: "app_id",
			Name:          "Test Service",
			Description:   "Description",
			Endpoint:      "https://service.daystram.com",
			Type:          "http",
			Config:        "{}",
			Showcase:      true,
		}).Return("", errors.New(""))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/application/app_id/service/", strings.NewReader(`{"name":"Test Service", "description": "Description", 
		"endpoint":"https://service.daystram.com", "type":"http", "config":"{}", "showcase":true}`))
		req.Header.Set("Content-Type", "application/json")
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Contains(suite.T(), w.Body.String(), `"error"`)
	})
}

func (suite *V1ServiceTestSuite) TestGETServiceDetail() {
	suite.T().Run("application exists", func(t *testing.T) {
		suite.MockHandler.EXPECT().ServiceGetOne("service_id", "app_id").Return(datatransfers.ServiceInfo{
			ID:            "service_id",
			ApplicationID: "app_id",
			Name:          "Test Service",
			Description:   "Description",
			Endpoint:      "https://service.daystram.com",
			Type:          "http",
			Config:        "{}",
			Showcase:      true,
			CreatedAt:     1,
			UpdatedAt:     2,
		}, nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/application/app_id/service/service_id", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.JSONEq(suite.T(), `{"data":{"id":"service_id", "name":"Test Service", "description": "Description", 
		"endpoint":"https://service.daystram.com", "type":"http", "config":"{}", "showcase":true, "createdAt": 1, "updatedAt": 2}}`, w.Body.String())
	})
	suite.T().Run("no application", func(t *testing.T) {
		suite.MockHandler.EXPECT().ServiceGetOne("service_id", "app_id").Return(datatransfers.ServiceInfo{}, gorm.ErrRecordNotFound)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/application/app_id/service/service_id", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusNotFound, w.Code)
		assert.JSONEq(suite.T(), `{}`, w.Body.String())
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockHandler.EXPECT().ServiceGetOne("service_id", "app_id").Return(datatransfers.ServiceInfo{}, errors.New(""))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/application/app_id/service/service_id", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Contains(suite.T(), w.Body.String(), `"error"`)
	})
}

func (suite *V1ServiceTestSuite) TestPUTServiceUpdate() {
	suite.T().Run("update", func(t *testing.T) {
		suite.MockHandler.EXPECT().ServiceUpdate(datatransfers.ServiceInfo{
			ID:            "service_id",
			ApplicationID: "app_id",
			Name:          "Test Service",
			Description:   "Description",
			Endpoint:      "https://service.daystram.com",
			Type:          "http",
			Config:        "{}",
			Showcase:      true,
		}).Return(nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/application/app_id/service/service_id", strings.NewReader(`{"name":"Test Service", "description": "Description",
		"endpoint":"https://service.daystram.com", "type":"http", "config":"{}", "showcase":true}`))
		req.Header.Set("Content-Type", "application/json")
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.JSONEq(suite.T(), `{}`, w.Body.String())
	})
	suite.T().Run("bad request", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/application/app_id/service/service_id", strings.NewReader(`{"description": "Description"}`))
		req.Header.Set("Content-Type", "application/json")
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	})
	suite.T().Run("invalid type", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/application/app_id/service/service_id", strings.NewReader(`{"name":"Test Service", "description": "Description", 
		"endpoint":"https://service.daystram.com", "type":"whattypeisthis", "config":"{}", "showcase":true}`))
		req.Header.Set("Content-Type", "application/json")
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockHandler.EXPECT().ServiceUpdate(datatransfers.ServiceInfo{
			ID:            "service_id",
			ApplicationID: "app_id",
			Name:          "Test Service",
			Description:   "Description",
			Endpoint:      "https://service.daystram.com",
			Type:          "http",
			Config:        "{}",
			Showcase:      true,
		}).Return(errors.New(""))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/application/app_id/service/service_id", strings.NewReader(`{"name":"Test Service", "description": "Description", 
		"endpoint":"https://service.daystram.com", "type":"http", "config":"{}", "showcase":true}`))
		req.Header.Set("Content-Type", "application/json")
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Contains(suite.T(), w.Body.String(), `"error"`)
	})
}

func (suite *V1ServiceTestSuite) TestDELETEServiceDelete() {
	suite.T().Run("delete", func(t *testing.T) {
		suite.MockHandler.EXPECT().ServiceDelete("service_id", "app_id").Return(nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/application/app_id/service/service_id", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.JSONEq(suite.T(), `{}`, w.Body.String())
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockHandler.EXPECT().ServiceDelete("service_id", "app_id").Return(errors.New(""))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/application/app_id/service/service_id", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Contains(suite.T(), w.Body.String(), `"error"`)
	})
}

func TestV1ServiceTestSuite(t *testing.T) {
	suite.Run(t, new(V1ServiceTestSuite))
}
