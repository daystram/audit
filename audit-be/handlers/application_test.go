package handlers

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/daystram/audit/audit-be/datatransfers"
	mock_models "github.com/daystram/audit/audit-be/mocks/models"
	"github.com/daystram/audit/audit-be/models"
)

type HandlerApplicationTestSuite struct {
	suite.Suite
	Handler              *module
	MockApplicationOrmer *mock_models.MockApplicationOrmer
}

func (suite *HandlerApplicationTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.MockApplicationOrmer = mock_models.NewMockApplicationOrmer(ctrl)
	suite.Handler = &module{
		db: &dbEntity{
			applicationOrmer: suite.MockApplicationOrmer,
		},
	}
}

func (suite *HandlerApplicationTestSuite) TestApplicationGetAll() {
	suite.T().Run("applications exist", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().GetAll().Return(make([]models.Application, 10), nil)
		applications, err := suite.Handler.ApplicationGetAll()
		assert.Equal(t, 10, len(applications))
		assert.Nil(t, err)
	})
	suite.T().Run("no applications", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().GetAll().Return(make([]models.Application, 0), gorm.ErrRecordNotFound)
		applications, err := suite.Handler.ApplicationGetAll()
		assert.Equal(t, 0, len(applications))
		assert.Nil(t, err)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().GetAll().Return(nil, errors.New(""))
		applications, err := suite.Handler.ApplicationGetAll()
		assert.Equal(t, 0, len(applications))
		assert.NotNil(t, err)
	})
}

func (suite *HandlerApplicationTestSuite) TestApplicationGetOne() {
	suite.T().Run("application exists", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().GetOneByID(gomock.Eq("app_id")).Return(models.Application{
			ID:   "app_id",
			Name: "Test App",
		}, nil)
		application, err := suite.Handler.ApplicationGetOne("app_id")
		assert.Equal(t, "Test App", application.Name)
		assert.Nil(t, err)
	})
	suite.T().Run("no application", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().GetOneByID(gomock.Eq("app_id")).Return(models.Application{}, gorm.ErrRecordNotFound)
		application, err := suite.Handler.ApplicationGetOne("app_id")
		assert.Empty(t, application)
		assert.NotNil(t, err)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().GetOneByID(gomock.Eq("app_id")).Return(models.Application{}, errors.New(""))
		applications, err := suite.Handler.ApplicationGetOne("app_id")
		assert.Empty(t, applications)
		assert.NotNil(t, err)
	})
}

func (suite *HandlerApplicationTestSuite) TestApplicationCreate() {
	suite.T().Run("create", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().Insert(gomock.Any()).Return("app_id", nil)
		applicationID, err := suite.Handler.ApplicationCreate(datatransfers.ApplicationInfo{
			Name:        "Test App",
			Description: "Description",
		})
		assert.Equal(t, "app_id", applicationID)
		assert.Nil(t, err)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().Insert(gomock.Any()).Return("", errors.New(""))
		applicationID, err := suite.Handler.ApplicationCreate(datatransfers.ApplicationInfo{
			Name:        "Test App",
			Description: "Description",
		})
		assert.Empty(t, applicationID)
		assert.NotNil(t, err)
	})
}

func (suite *HandlerApplicationTestSuite) TestApplicationUpdate() {
	suite.T().Run("create", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().Update(gomock.Any()).Return(nil)
		err := suite.Handler.ApplicationUpdate(datatransfers.ApplicationInfo{
			Name:        "Test App",
			Description: "Description",
		})
		assert.Nil(t, err)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().Update(gomock.Any()).Return(errors.New(""))
		err := suite.Handler.ApplicationUpdate(datatransfers.ApplicationInfo{
			Name:        "Test App",
			Description: "Description",
		})
		assert.NotNil(t, err)
	})
}

func (suite *HandlerApplicationTestSuite) TestApplicationDelete() {
	suite.T().Run("delete", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().DeleteByID("app_id").Return(nil)
		err := suite.Handler.ApplicationDelete("app_id")
		assert.Nil(t, err)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().DeleteByID(gomock.Any()).Return(errors.New(""))
		err := suite.Handler.ApplicationDelete("app_id")
		assert.NotNil(t, err)
	})
}

func TestHandlerApplicationTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerApplicationTestSuite))
}
