package handlers

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	mock_models "github.com/daystram/audit/audit-be/mocks/models"
	"github.com/daystram/audit/audit-be/models"
)

type HandlerSchedulerTestSuite struct {
	suite.Suite
	Handler          *module
	MockServiceOrmer *mock_models.MockServiceOrmer
}

func (suite *HandlerSchedulerTestSuite) SetupTest(trackers map[string]TrackerClient) {
	ctrl := gomock.NewController(suite.T())
	suite.MockServiceOrmer = mock_models.NewMockServiceOrmer(ctrl)
	suite.Handler = &module{
		db: &dbEntity{
			serviceOrmer: suite.MockServiceOrmer,
		},
	}
}

func (suite *HandlerSchedulerTestSuite) TestInitializeScheduler() {
	suite.T().Run("starts", func(t *testing.T) {
		suite.SetupTest(make(map[string]TrackerClient))
		err := suite.Handler.InitializeScheduler()
		assert.Nil(suite.T(), err)
	})
}

func (suite *HandlerSchedulerTestSuite) TestTriggerTracking() {
	// TODO: "has service" test. error on MockTrackerServer not implementing interface from proto's mock
	suite.T().Run("no services", func(t *testing.T) {
		services := make(map[string]TrackerClient)
		suite.SetupTest(services)
		suite.MockServiceOrmer.EXPECT().GetAllEnabled().Return(make([]models.Service, 0), nil)
		suite.Handler.TriggerTracking()
	})
}

func TestHandlerSchedulerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerSchedulerTestSuite))
}
