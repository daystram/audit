package handlers

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/daystram/audit/audit-be/config"
	pb "github.com/daystram/audit/proto"
)

type TrackerServer struct {
	pb.UnimplementedTrackerServer
	trackers map[string]*TrackerClient
	mu       sync.RWMutex
}

type TrackerClient struct {
	stream     pb.Tracker_SubscribeServer
	lastPinged int64
	latency    int64
}

func (m *module) InitializeTrackerServer() {
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", config.AppConfig.TrackerServerPort))
	if err != nil {
		log.Fatalf("[TrackerServer] initialization failed. %v", err)
	}
	m.trackerServer = &TrackerServer{trackers: make(map[string]*TrackerClient)}
	// TODO: authentication
	pb.RegisterTrackerServer(grpcServer, m.trackerServer)
	go grpcServer.Serve(lis)
}

func (s *TrackerServer) Subscribe(request *pb.SubscriptionRequest, stream pb.Tracker_SubscribeServer) (err error) {
	trackerID := request.TrackerId
	// TODO: validate client
	s.mu.Lock()
	if _, ok := s.trackers[trackerID]; ok {
		s.mu.Unlock()
		return fmt.Errorf("tracker with ID %s already registered", trackerID)
	}
	s.trackers[trackerID] = &TrackerClient{
		stream: stream,
	}
	s.mu.Unlock()
	log.Printf("[TrackerServer] tracker %s connected\n", trackerID)
	s.SendTrackingRequest()
	s.pingTracker(trackerID) // keep alive
	return
}

func (s *TrackerServer) pingTracker(trackerID string) {
	s.mu.RLock()
	tracker := s.trackers[trackerID]
	s.mu.RUnlock()
	for {
		err := tracker.stream.Send(&pb.TrackingMessage{
			Code: pb.MessageType_MESSAGE_TYPE_PING,
			Body: &pb.TrackingMessage_Request{
				Request: &pb.TrackingRequest{
					TrackerId:   trackerID,
					RequestedAt: time.Now().UnixNano(), // only PING uses Unix nano
				},
			},
		})
		if err != nil {
			s.mu.Lock()
			delete(s.trackers, trackerID)
			s.mu.Unlock()
			log.Printf("[TrackerServer] %s disconnected. remaining trackers: %d\n", trackerID, len(s.trackers))
			return
		}
		time.Sleep(time.Second)
	}
}

func (s *TrackerServer) SendTrackingRequest() {
	// TODO: example setup; implement
	s.mu.Lock()
	defer s.mu.Unlock()
	for trackerID, tracker := range s.trackers {
		tracker.stream.Send(&pb.TrackingMessage{
			Code: pb.MessageType_MESSAGE_TYPE_TRACKING,
			Body: &pb.TrackingMessage_Request{
				Request: &pb.TrackingRequest{
					ApplicationId: "app_id",
					ServiceId:     "service_id",
					TrackerId:     trackerID,
					Endpoint:      "service.daystram.com:80",
				},
			},
		})
		log.Printf("[TrackerServer] send request to %s", trackerID)
	}
}

func (s *TrackerServer) ReportTrackingRequest(ctx context.Context, message *pb.TrackingMessage) (*pb.Empty, error) {
	// TODO: implement
	response := message.Body.(*pb.TrackingMessage_Response).Response
	log.Printf("[TrackerServer] receive tracking response from %s: %+v", response.TrackerId, response)
	return &pb.Empty{}, nil
}

func (s *TrackerServer) Pong(ctx context.Context, message *pb.TrackingMessage) (*pb.Empty, error) {
	request := message.Body.(*pb.TrackingMessage_Request).Request
	trackerID := request.TrackerId
	latency := time.Now().UnixNano() - request.RequestedAt
	log.Printf("[TrackerServer] latency to %s: %dms", trackerID, latency/10e6)
	s.mu.RLock()
	s.trackers[trackerID].lastPinged = time.Now().Unix()
	s.trackers[trackerID].latency = latency
	s.mu.RUnlock()
	return &pb.Empty{}, nil
}
