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

type V1ApplicationTestSuite struct {
	suite.Suite
	Router      *gin.Engine
	MockHandler *mock_handlers.MockHandlerFunc
}

func (suite *V1ApplicationTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.MockHandler = mock_handlers.NewMockHandlerFunc(ctrl)
	suite.Router = InitializeRouter(suite.MockHandler)
	gin.SetMode(gin.ReleaseMode)
}

func (suite *V1ApplicationTestSuite) TestGETApplicationList() {
	suite.T().Run("applications exist", func(t *testing.T) {
		suite.MockHandler.EXPECT().ApplicationGetAll().Return([]datatransfers.ApplicationInfo{{
			ID:          "app_id",
			Name:        "Test App",
			Description: "Description",
			CreatedAt:   1,
			UpdatedAt:   2,
		}}, nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/application/", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.JSONEq(suite.T(), `{"data": [{"id":"app_id", "name":"Test App", "description": "Description", "createdAt": 1, "updatedAt": 2}]}`, w.Body.String())
	})
	suite.T().Run("no applications", func(t *testing.T) {
		suite.MockHandler.EXPECT().ApplicationGetAll().Return(make([]datatransfers.ApplicationInfo, 0), nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/application/", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.JSONEq(suite.T(), `{"data": []}`, w.Body.String())
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockHandler.EXPECT().ApplicationGetAll().Return([]datatransfers.ApplicationInfo{}, errors.New(""))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/application/", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Contains(suite.T(), w.Body.String(), `"error"`)
	})
}

func (suite *V1ApplicationTestSuite) TestPOSTApplicationCreate() {
	suite.T().Run("create", func(t *testing.T) {
		suite.MockHandler.EXPECT().ApplicationCreate(datatransfers.ApplicationInfo{
			Name:        "Test App",
			Description: "Description",
		}).Return("app_id", nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/application/", strings.NewReader(`{"name": "Test App", "description": "Description"}`))
		req.Header.Set("Content-Type", "application/json")
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.JSONEq(suite.T(), `{}`, w.Body.String())
	})
	suite.T().Run("bad request", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/application/", strings.NewReader(`{"description": "Description"}`))
		req.Header.Set("Content-Type", "application/json")
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockHandler.EXPECT().ApplicationCreate(datatransfers.ApplicationInfo{
			Name:        "Test App",
			Description: "Description",
		}).Return("", errors.New(""))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/application/", strings.NewReader(`{"name": "Test App", "description": "Description"}`))
		req.Header.Set("Content-Type", "application/json")
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Contains(suite.T(), w.Body.String(), `"error"`)
	})
}

func (suite *V1ApplicationTestSuite) TestGETApplicationDetail() {
	suite.T().Run("application exists", func(t *testing.T) {
		suite.MockHandler.EXPECT().ApplicationGetOne("app_id").Return(datatransfers.ApplicationInfo{
			ID:          "app_id",
			Name:        "Test App",
			Description: "Description",
			CreatedAt:   1,
			UpdatedAt:   2,
		}, nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/application/app_id", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.JSONEq(suite.T(), `{"data": {"id":"app_id", "name":"Test App", "description": "Description", "createdAt": 1, "updatedAt": 2}}`, w.Body.String())
	})
	suite.T().Run("no application", func(t *testing.T) {
		suite.MockHandler.EXPECT().ApplicationGetOne("app_id").Return(datatransfers.ApplicationInfo{}, gorm.ErrRecordNotFound)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/application/app_id", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusNotFound, w.Code)
		assert.JSONEq(suite.T(), `{}`, w.Body.String())
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockHandler.EXPECT().ApplicationGetOne("app_id").Return(datatransfers.ApplicationInfo{}, errors.New(""))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/application/app_id", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Contains(suite.T(), w.Body.String(), `"error"`)
	})
}

func (suite *V1ApplicationTestSuite) TestPUTApplicationUpdate() {
	suite.T().Run("update", func(t *testing.T) {
		suite.MockHandler.EXPECT().ApplicationUpdate(datatransfers.ApplicationInfo{
			ID:          "app_id",
			Name:        "Test App",
			Description: "Description",
		}).Return(nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/api/application/app_id", strings.NewReader(`{"name": "Test App", "description": "Description"}`))
		req.Header.Set("Content-Type", "application/json")
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.JSONEq(suite.T(), `{}`, w.Body.String())
	})
	suite.T().Run("bad request", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/api/application/app_id", strings.NewReader(`{"description": "Description"}`))
		req.Header.Set("Content-Type", "application/json")
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockHandler.EXPECT().ApplicationUpdate(datatransfers.ApplicationInfo{
			ID:          "app_id",
			Name:        "Test App",
			Description: "Description",
		}).Return(errors.New(""))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPut, "/api/application/app_id", strings.NewReader(`{"name": "Test App", "description": "Description"}`))
		req.Header.Set("Content-Type", "application/json")
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Contains(suite.T(), w.Body.String(), `"error"`)
	})
}

func (suite *V1ApplicationTestSuite) TestDELETEApplicationDelete() {
	suite.T().Run("delete", func(t *testing.T) {
		suite.MockHandler.EXPECT().ApplicationDelete("app_id").Return(nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/api/application/app_id", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusOK, w.Code)
		assert.JSONEq(suite.T(), `{}`, w.Body.String())
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockHandler.EXPECT().ApplicationDelete("app_id").Return(errors.New(""))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/api/application/app_id", nil)
		suite.Router.ServeHTTP(w, req)
		assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
		assert.Contains(suite.T(), w.Body.String(), `"error"`)
	})
}

func TestV1ApplicationTestSuite(t *testing.T) {
	suite.Run(t, new(V1ApplicationTestSuite))
}
