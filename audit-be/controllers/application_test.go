package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type V1ApplicationTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *V1ApplicationTestSuite) SetupTest() {
	suite.router = InitializeRouter()
}

func (suite *V1ApplicationTestSuite) TestGETApplicationList() {
	suite.T().Run("applications exist", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ping", nil)
		suite.router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "pong", w.Body.String())
	})
}

func TestV1ApplicationTestSuite(t *testing.T) {
	suite.Run(t, new(V1ApplicationTestSuite))
}
