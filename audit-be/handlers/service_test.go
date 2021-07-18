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

type HandlerServiceTestSuite struct {
	suite.Suite
	Handler          *module
	MockServiceOrmer *mock_models.MockServiceOrmer
}

func (suite *HandlerServiceTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.MockServiceOrmer = mock_models.NewMockServiceOrmer(ctrl)
	suite.Handler = &module{
		db: &dbEntity{
			serviceOrmer: suite.MockServiceOrmer,
		},
	}
}

func (suite *HandlerServiceTestSuite) TestServiceGetAll() {
	suite.T().Run("services exist", func(t *testing.T) {
		suite.MockServiceOrmer.EXPECT().GetAllByApplicationID("app_id").Return(make([]models.Service, 10), nil)
		service, err := suite.Handler.ServiceGetAll("app_id")
		assert.Equal(t, 10, len(service))
		assert.Nil(t, err)
	})
	suite.T().Run("no services", func(t *testing.T) {
		suite.MockServiceOrmer.EXPECT().GetAllByApplicationID("app_id").Return(make([]models.Service, 0), gorm.ErrRecordNotFound)
		service, err := suite.Handler.ServiceGetAll("app_id")
		assert.Equal(t, 0, len(service))
		assert.Nil(t, err)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockServiceOrmer.EXPECT().GetAllByApplicationID("app_id").Return(nil, errors.New(""))
		service, err := suite.Handler.ServiceGetAll("app_id")
		assert.Equal(t, 0, len(service))
		assert.NotNil(t, err)
	})
}

func (suite *HandlerServiceTestSuite) TestServiceGetOne() {
	suite.T().Run("service exists", func(t *testing.T) {
		suite.MockServiceOrmer.EXPECT().GetOneByIDAndApplicationID("service_id", "app_id").Return(models.Service{
			ID:            "service_id",
			ApplicationID: "app_id",
			Name:          "Test Service",
		}, nil)
		service, err := suite.Handler.ServiceGetOne("service_id", "app_id")
		assert.Equal(t, "Test Service", service.Name)
		assert.Nil(t, err)
	})
	suite.T().Run("no service", func(t *testing.T) {
		suite.MockServiceOrmer.EXPECT().GetOneByIDAndApplicationID("service_id", "app_id").Return(models.Service{}, gorm.ErrRecordNotFound)
		service, err := suite.Handler.ServiceGetOne("service_id", "app_id")
		assert.Empty(t, service)
		assert.NotNil(t, err)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockServiceOrmer.EXPECT().GetOneByIDAndApplicationID("service_id", "app_id").Return(models.Service{}, errors.New(""))
		service, err := suite.Handler.ServiceGetOne("service_id", "app_id")
		assert.Empty(t, service)
		assert.NotNil(t, err)
	})
}

func (suite *HandlerServiceTestSuite) TestServiceCreate() {
	suite.T().Run("create", func(t *testing.T) {
		suite.MockServiceOrmer.EXPECT().Insert(gomock.Any()).Return("app_id", nil)
		serviceID, err := suite.Handler.ServiceCreate(datatransfers.ServiceInfo{
			ApplicationID: "app_id",
			Name:          "Test Service",
			Description:   "Description",
		})
		assert.Equal(t, "app_id", serviceID)
		assert.Nil(t, err)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockServiceOrmer.EXPECT().Insert(gomock.Any()).Return("", errors.New(""))
		serviceID, err := suite.Handler.ServiceCreate(datatransfers.ServiceInfo{
			ApplicationID: "app_id",
			Name:          "Test Service",
			Description:   "Description",
		})
		assert.Empty(t, serviceID)
		assert.NotNil(t, err)
	})
}

func (suite *HandlerServiceTestSuite) TestServiceUpdate() {
	suite.T().Run("create", func(t *testing.T) {
		suite.MockServiceOrmer.EXPECT().Update(gomock.Any()).Return(nil)
		err := suite.Handler.ServiceUpdate(datatransfers.ServiceInfo{
			ID:            "service_id",
			ApplicationID: "app_id",
			Name:          "Test Service",
			Description:   "Description",
		})
		assert.Nil(t, err)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockServiceOrmer.EXPECT().Update(gomock.Any()).Return(errors.New(""))
		err := suite.Handler.ServiceUpdate(datatransfers.ServiceInfo{
			ID:            "service_id",
			ApplicationID: "app_id",
			Name:          "Test Service",
			Description:   "Description",
		})
		assert.NotNil(t, err)
	})
}

func (suite *HandlerServiceTestSuite) TestServiceDelete() {
	suite.T().Run("delete", func(t *testing.T) {
		suite.MockServiceOrmer.EXPECT().DeleteByIDAndApplicationID("service_id", "app_id").Return(nil)
		err := suite.Handler.ServiceDelete("service_id", "app_id")
		assert.Nil(t, err)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockServiceOrmer.EXPECT().DeleteByIDAndApplicationID("service_id", "app_id").Return(errors.New(""))
		err := suite.Handler.ServiceDelete("service_id", "app_id")
		assert.NotNil(t, err)
	})
}

func TestHandlerServiceTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerServiceTestSuite))
}
