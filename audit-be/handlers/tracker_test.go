package handlers

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/daystram/audit/audit-be/constants"
	mock_models "github.com/daystram/audit/audit-be/mocks/models"
	"github.com/daystram/audit/audit-be/models"
	pb "github.com/daystram/audit/proto"
	mock_pb "github.com/daystram/audit/proto/mocks"
)

type HandlerTrackerTestSuite struct {
	suite.Suite
	Handler *module
}

func (suite *HandlerTrackerTestSuite) SetupTest(trackers map[string]TrackerClient) {
	suite.Handler = &module{
		trackerServer: &trackerServerModule{
			trackers: trackers,
		},
	}
}

func (suite *HandlerTrackerTestSuite) SetupTestWithID(trackers map[string]TrackerClient, trackerIDs []string) {
	suite.Handler = &module{
		trackerServer: &trackerServerModule{
			trackers:   trackers,
			trackerIDs: trackerIDs,
		},
	}
}

func (suite *HandlerTrackerTestSuite) SetupTestWithInjectedHandler(trackers map[string]TrackerClient, handlers *module) {
	suite.Handler = &module{
		trackerServer: &trackerServerModule{
			handlers: handlers,
			trackers: trackers,
		},
	}
}

func (suite *HandlerTrackerTestSuite) TestInitializeTrackerServer() {
	suite.T().Run("starts", func(t *testing.T) {
		suite.SetupTest(make(map[string]TrackerClient))
		err := suite.Handler.InitializeTrackerServer(7899)
		assert.Nil(suite.T(), err)
	})
}

func (suite *HandlerTrackerTestSuite) TestSubscribe() {
	suite.T().Run("registers", func(t *testing.T) {
		trackers := make(map[string]TrackerClient)
		suite.SetupTest(trackers)
		ctrl := gomock.NewController(suite.T())
		defer ctrl.Finish()
		mockStream := mock_pb.NewMockTracker_SubscribeServer(ctrl)
		mockStream.EXPECT().Send(gomock.AssignableToTypeOf(&pb.TrackingMessage{})).Times(1).Return(nil)
		mockStream.EXPECT().Send(gomock.AssignableToTypeOf(&pb.TrackingMessage{})).Times(1).Return(errors.New(""))
		err := suite.Handler.trackerServer.Subscribe(&pb.SubscriptionRequest{
			TrackerId: "tracker_id",
		}, mockStream)
		assert.Empty(suite.T(), trackers)
		assert.NotNil(suite.T(), err)
	})
	suite.T().Run("already exists", func(t *testing.T) {
		trackers := make(map[string]TrackerClient)
		suite.SetupTest(trackers)
		ctrl := gomock.NewController(suite.T())
		defer ctrl.Finish()
		trackers["tracker_id"] = &trackerClientEntity{}
		mockStream := mock_pb.NewMockTracker_SubscribeServer(ctrl)
		err := suite.Handler.trackerServer.Subscribe(&pb.SubscriptionRequest{
			TrackerId: "tracker_id",
		}, mockStream)
		assert.Len(suite.T(), trackers, 1)
		assert.NotNil(suite.T(), err)
	})
}

func (suite *HandlerTrackerTestSuite) TestSendTrackingRequest() {
	suite.T().Run("runs", func(t *testing.T) {
		trackers := make(map[string]TrackerClient)
		trackerIDs := []string{"tracker_id"}
		suite.SetupTestWithID(trackers, trackerIDs)
		ctrl := gomock.NewController(suite.T())
		defer ctrl.Finish()
		mockStream := mock_pb.NewMockTracker_SubscribeServer(ctrl)
		mockStream.EXPECT().Send(&pb.TrackingMessage{
			Code: pb.MessageType_MESSAGE_TYPE_TRACKING,
			Body: &pb.TrackingMessage_Request{
				Request: &pb.TrackingRequest{
					ApplicationId: "app_id",
					ServiceId:     "service_id",
					TrackerId:     "tracker_id",
					Endpoint:      "service.daystram.com:80",
				},
			},
		}).Times(1)
		trackers["tracker_id"] = &trackerClientEntity{
			stream: mockStream,
		}
		err := suite.Handler.trackerServer.SendTrackingRequest(&pb.TrackingRequest{
			ApplicationId: "app_id",
			ServiceId:     "service_id",
			Endpoint:      "service.daystram.com:80",
		})
		assert.Nil(suite.T(), err)
	})
	suite.T().Run("unregistered", func(t *testing.T) {
		trackers := make(map[string]TrackerClient)
		trackerIDs := make([]string, 0)
		suite.SetupTestWithID(trackers, trackerIDs)
		err := suite.Handler.trackerServer.SendTrackingRequest(&pb.TrackingRequest{
			ApplicationId: "app_id",
			ServiceId:     "service_id",
			TrackerId:     "tracker_id",
			Endpoint:      "service.daystram.com:80",
		})
		assert.NotNil(suite.T(), err)
	})
}

func (suite *HandlerTrackerTestSuite) TestReportTrackingRequest() {
	suite.T().Run("HTTP service", func(t *testing.T) {
		trackers := make(map[string]TrackerClient)
		ctrl := gomock.NewController(suite.T())
		defer ctrl.Finish()
		mockApplicationOrmer := mock_models.NewMockApplicationOrmer(ctrl)
		mockServiceOrmer := mock_models.NewMockServiceOrmer(ctrl)
		handlers := &module{
			db: &dbEntity{
				applicationOrmer: mockApplicationOrmer,
				serviceOrmer:     mockServiceOrmer,
			},
		}
		suite.SetupTestWithInjectedHandler(trackers, handlers)
		mockApplicationOrmer.EXPECT().GetOneByID("app_id").Return(models.Application{ID: "app_id"}, nil)
		mockServiceOrmer.EXPECT().GetOneByIDAndApplicationID("service_id", "app_id").Return(models.Service{ID: "service_id", Type: constants.ServiceTypeHTTP}, nil)
		_, err := suite.Handler.trackerServer.ReportTrackingRequest(context.Background(), &pb.TrackingMessage{
			Code: pb.MessageType_MESSAGE_TYPE_TRACKING,
			Body: &pb.TrackingMessage_Response{
				Response: &pb.TrackingResponse{
					ApplicationId: "app_id",
					ServiceId:     "service_id",
					TrackerId:     "tracker_id",
					Status:        pb.ServiceStatus_SERVICE_STATUS_UP,
					Body:          "200",
					ResponseTime:  (1 * time.Second).Nanoseconds(),
				},
			},
		})
		assert.Nil(suite.T(), err)
	})
	suite.T().Run("TCP service", func(t *testing.T) {
		trackers := make(map[string]TrackerClient)
		ctrl := gomock.NewController(suite.T())
		defer ctrl.Finish()
		mockApplicationOrmer := mock_models.NewMockApplicationOrmer(ctrl)
		mockServiceOrmer := mock_models.NewMockServiceOrmer(ctrl)
		handlers := &module{
			db: &dbEntity{
				applicationOrmer: mockApplicationOrmer,
				serviceOrmer:     mockServiceOrmer,
			},
		}
		suite.SetupTestWithInjectedHandler(trackers, handlers)
		mockApplicationOrmer.EXPECT().GetOneByID("app_id").Return(models.Application{ID: "app_id"}, nil)
		mockServiceOrmer.EXPECT().GetOneByIDAndApplicationID("service_id", "app_id").Return(models.Service{ID: "service_id", Type: constants.ServiceTypeTCP}, nil)
		_, err := suite.Handler.trackerServer.ReportTrackingRequest(context.Background(), &pb.TrackingMessage{
			Code: pb.MessageType_MESSAGE_TYPE_TRACKING,
			Body: &pb.TrackingMessage_Response{
				Response: &pb.TrackingResponse{
					ApplicationId: "app_id",
					ServiceId:     "service_id",
					TrackerId:     "tracker_id",
					Status:        pb.ServiceStatus_SERVICE_STATUS_UP,
					ResponseTime:  (1 * time.Second).Nanoseconds(),
				},
			},
		})
		assert.Nil(suite.T(), err)
	})
	suite.T().Run("PING service", func(t *testing.T) {
		trackers := make(map[string]TrackerClient)
		ctrl := gomock.NewController(suite.T())
		defer ctrl.Finish()
		mockApplicationOrmer := mock_models.NewMockApplicationOrmer(ctrl)
		mockServiceOrmer := mock_models.NewMockServiceOrmer(ctrl)
		handlers := &module{
			db: &dbEntity{
				applicationOrmer: mockApplicationOrmer,
				serviceOrmer:     mockServiceOrmer,
			},
		}
		suite.SetupTestWithInjectedHandler(trackers, handlers)
		mockApplicationOrmer.EXPECT().GetOneByID("app_id").Return(models.Application{ID: "app_id"}, nil)
		mockServiceOrmer.EXPECT().GetOneByIDAndApplicationID("service_id", "app_id").Return(models.Service{ID: "service_id", Type: constants.ServiceTypePING}, nil)
		_, err := suite.Handler.trackerServer.ReportTrackingRequest(context.Background(), &pb.TrackingMessage{
			Code: pb.MessageType_MESSAGE_TYPE_TRACKING,
			Body: &pb.TrackingMessage_Response{
				Response: &pb.TrackingResponse{
					ApplicationId: "app_id",
					ServiceId:     "service_id",
					TrackerId:     "tracker_id",
					Status:        pb.ServiceStatus_SERVICE_STATUS_UP,
					ResponseTime:  (1 * time.Second).Nanoseconds(),
				},
			},
		})
		assert.Nil(suite.T(), err)
	})
}

func (suite *HandlerTrackerTestSuite) TestPong() {
	suite.T().Run("runs", func(t *testing.T) {
		trackers := make(map[string]TrackerClient)
		suite.SetupTest(trackers)
		trackers["tracker_id"] = &trackerClientEntity{
			lastPinged: -1,
			latency:    -1,
		}
		_, err := suite.Handler.trackerServer.Pong(context.Background(), &pb.TrackingMessage{
			Code: pb.MessageType_MESSAGE_TYPE_PING,
			Body: &pb.TrackingMessage_Request{
				Request: &pb.TrackingRequest{
					TrackerId:   "tracker_id",
					RequestedAt: time.Now().Add(-time.Second).UnixNano(),
				},
			},
		})
		lastPinged, latency := trackers["tracker_id"].Status()
		assert.Nil(suite.T(), err)
		assert.NotEqual(suite.T(), -1, lastPinged)
		assert.InEpsilon(suite.T(), time.Second, latency, 1e-3)
	})
	suite.T().Run("unregistered", func(t *testing.T) {
		trackers := make(map[string]TrackerClient)
		suite.SetupTest(trackers)
		_, err := suite.Handler.trackerServer.Pong(context.Background(), &pb.TrackingMessage{
			Code: pb.MessageType_MESSAGE_TYPE_PING,
			Body: &pb.TrackingMessage_Request{
				Request: &pb.TrackingRequest{
					TrackerId:   "tracker_id",
					RequestedAt: time.Now().Add(-time.Second).UnixNano(),
				},
			},
		})
		assert.NotNil(suite.T(), err)
	})
}

func TestHandlerTrackerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTrackerTestSuite))
}
