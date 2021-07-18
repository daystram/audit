package handlers

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	pb "github.com/daystram/audit/proto"
	mock_pb "github.com/daystram/audit/proto/mocks"
)

type HandlerTrackerTestSuite struct {
	suite.Suite
	Handler *module
}

func (suite *HandlerTrackerTestSuite) SetupTest() {
	suite.Handler = &module{
		trackerServer: &TrackerServer{
			trackers: map[string]*TrackerClient{},
			mu:       sync.RWMutex{},
		},
	}
}

func (suite *HandlerTrackerTestSuite) TestInitializeTrackerServer() {
	suite.T().Run("starts", func(t *testing.T) {
		suite.Handler.InitializeTrackerServer(9999)
	})
}

func (suite *HandlerTrackerTestSuite) TestSubscribe() {
	suite.T().Run("registers", func(t *testing.T) {
		ctrl := gomock.NewController(suite.T())
		defer ctrl.Finish()
		mockStream := mock_pb.NewMockTracker_SubscribeServer(ctrl)
		mockStream.EXPECT().Send(gomock.AssignableToTypeOf(&pb.TrackingMessage{})).Times(1).Return(nil)
		mockStream.EXPECT().Send(gomock.AssignableToTypeOf(&pb.TrackingMessage{})).Times(1).Return(errors.New(""))
		err := suite.Handler.trackerServer.Subscribe(&pb.SubscriptionRequest{
			TrackerId: "tracker_id",
		}, mockStream)
		assert.Empty(suite.T(), suite.Handler.trackerServer.trackers)
		assert.NotNil(suite.T(), err)
	})
	suite.T().Run("already exists", func(t *testing.T) {
		ctrl := gomock.NewController(suite.T())
		defer ctrl.Finish()
		suite.Handler.trackerServer.trackers["tracker_id"] = &TrackerClient{}
		mockStream := mock_pb.NewMockTracker_SubscribeServer(ctrl)
		err := suite.Handler.trackerServer.Subscribe(&pb.SubscriptionRequest{
			TrackerId: "tracker_id",
		}, mockStream)
		assert.Len(suite.T(), suite.Handler.trackerServer.trackers, 1)
		assert.NotNil(suite.T(), err)
	})
}

func (suite *HandlerTrackerTestSuite) TestSendTrackingRequest() {
	suite.T().Run("runs", func(t *testing.T) {
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
		suite.Handler.trackerServer.trackers["tracker_id"] = &TrackerClient{
			stream: mockStream,
		}
		suite.Handler.trackerServer.SendTrackingRequest(&pb.TrackingRequest{
			ApplicationId: "app_id",
			ServiceId:     "service_id",
			TrackerId:     "tracker_id",
			Endpoint:      "service.daystram.com:80",
		})
	})
}

func (suite *HandlerTrackerTestSuite) TestReportTrackingRequest() {
	suite.T().Run("runs", func(t *testing.T) {
		suite.Handler.trackerServer.trackers["tracker_id"] = &TrackerClient{}
		suite.Handler.trackerServer.ReportTrackingRequest(context.Background(), &pb.TrackingMessage{
			Code: pb.MessageType_MESSAGE_TYPE_TRACKING,
			Body: &pb.TrackingMessage_Response{
				Response: &pb.TrackingResponse{
					ApplicationId: "app_id",
					ServiceId:     "service_id",
					TrackerId:     "tracker_id",
				},
			},
		})
	})
}

func (suite *HandlerTrackerTestSuite) TestPong() {
	suite.T().Run("runs", func(t *testing.T) {
		now := time.Now().Add(-time.Second)
		suite.Handler.trackerServer.trackers["tracker_id"] = &TrackerClient{
			lastPinged: 0,
			latency:    0,
		}
		suite.Handler.trackerServer.Pong(context.Background(), &pb.TrackingMessage{
			Code: pb.MessageType_MESSAGE_TYPE_PING,
			Body: &pb.TrackingMessage_Request{
				Request: &pb.TrackingRequest{
					TrackerId:   "tracker_id",
					RequestedAt: now.UnixNano(),
				},
			},
		})
		assert.NotZero(suite.T(), suite.Handler.trackerServer.trackers["tracker_id"].lastPinged)
		assert.InEpsilon(suite.T(), time.Second, suite.Handler.trackerServer.trackers["tracker_id"].latency, 1e-3)
	})
}

func TestHandlerTrackerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTrackerTestSuite))
}
