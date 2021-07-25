package handlers

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	mock_models "github.com/daystram/audit/audit-be/mocks/models"
	"github.com/daystram/audit/audit-be/models"
)

type HandlerMonitorTestSuite struct {
	suite.Suite
	Handler              *module
	MockApplicationOrmer *mock_models.MockApplicationOrmer
	MockReportOrmer      *mock_models.MockReportOrmer
}

func (suite *HandlerMonitorTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.MockApplicationOrmer = mock_models.NewMockApplicationOrmer(ctrl)
	suite.MockReportOrmer = mock_models.NewMockReportOrmer(ctrl)
	suite.Handler = &module{
		db: &dbEntity{
			applicationOrmer: suite.MockApplicationOrmer,
		},
		influx: &influxEntity{
			reportOrmer: suite.MockReportOrmer,
		},
	}
}

func (suite *HandlerMonitorTestSuite) TestMonitorGetAll() {
	suite.T().Run("applications exist", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().GetAllShowcaseWithServices().Return([]models.Application{{
			ID:       "app_id",
			Services: []models.Service{{ID: "service_id1"}, {ID: "service_id2"}},
		}}, nil)
		suite.MockReportOrmer.EXPECT().GetWindowByApplicationIDAndServiceID("app_id", "service_id1").Return([]models.Report{{}}, nil)
		suite.MockReportOrmer.EXPECT().GetWindowByApplicationIDAndServiceID("app_id", "service_id2").Return([]models.Report{{}}, nil)
		applications, err := suite.Handler.MonitorGetAll()
		assert.Equal(t, 1, len(applications))
		assert.Equal(t, 2, len(applications[0].Services))
		assert.Nil(t, err)
	})
	suite.T().Run("no applications", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().GetAllShowcaseWithServices().Return(make([]models.Application, 0), gorm.ErrRecordNotFound)
		applications, err := suite.Handler.MonitorGetAll()
		assert.Equal(t, 0, len(applications))
		assert.Nil(t, err)
	})
	suite.T().Run("has error", func(t *testing.T) {
		suite.MockApplicationOrmer.EXPECT().GetAllShowcaseWithServices().Return(nil, errors.New(""))
		applications, err := suite.Handler.MonitorGetAll()
		assert.Equal(t, 0, len(applications))
		assert.NotNil(t, err)
	})
}

func TestHandlerMonitorTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerMonitorTestSuite))
}
