package handlers

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/daystram/audit/proto"
)

type TrackerServer struct {
	pb.UnimplementedTrackerServer
	// TODO: thread safety
	trackers map[string]pb.Tracker_SubscribeServer
}

func (m *module) InitializeTrackerServer() {
	grpcServer := grpc.NewServer()
	// TODO: use app config
	lis, err := net.Listen("tcp", "localhost:5555")
	if err != nil {
		log.Fatalf("[TrackerServer] failed initializing. %v", err)
	}
	m.trackerServer = &TrackerServer{trackers: make(map[string]pb.Tracker_SubscribeServer)}
	// TODO: authentication
	pb.RegisterTrackerServer(grpcServer, m.trackerServer)
	go grpcServer.Serve(lis)
}

func (s *TrackerServer) Subscribe(request *pb.SubscriptionRequest, stream pb.Tracker_SubscribeServer) (err error) {
	trackerID := request.TrackerId
	s.trackers[trackerID] = stream
	// TODO: validate client
	log.Printf("[TrackerServer] %s connected\n", trackerID)
	s.SendTrackingRequest()
	s.pingTracker(trackerID)
	return
}

func (s *TrackerServer) pingTracker(trackerID string) {
	// TODO: tracker latency
	stream := s.trackers[trackerID]
	for {
		d, _ := time.ParseDuration("1s")
		time.Sleep(d)
		err := stream.Send(&pb.TrackingMessage{
			Code: pb.MessageType_MESSAGE_TYPE_PING,
		})
		if err != nil {
			delete(s.trackers, trackerID)
			log.Printf("[TrackerServer] %s disconnected. remaining trackers: %d\n", trackerID, len(s.trackers))
			return
		}
	}
}

func (s *TrackerServer) SendTrackingRequest() {
	// TODO: example setup; implement
	for subscriberID, stream := range s.trackers {
		stream.Send(&pb.TrackingMessage{
			Code: pb.MessageType_MESSAGE_TYPE_TRACKING,
			Body: &pb.TrackingMessage_Request{
				Request: &pb.TrackingRequest{
					ApplicationId: "app_id",
					ServiceId:     "service_id",
					Endpoint:      "service.daystram.com:80",
				},
			},
		})
		log.Printf("[TrackerServer] send request to %s", subscriberID)
	}
}

func (s *TrackerServer) ReportTrackingRequest(ctx context.Context, message *pb.TrackingMessage) (_ *pb.Empty, err error) {
	// TODO: implement
	return
}

func (s *TrackerServer) Pong(ctx context.Context, message *pb.TrackingMessage) (_ *pb.Empty, err error) {
	// TODO: implement
	return
}
