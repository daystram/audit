package handlers

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/daystram/audit/audit-be/constants"
	pb "github.com/daystram/audit/proto"
)

type TrackerServer interface {
	pb.TrackerServer
	PingTracker(trackerID string) (err error)
	SendTrackingRequest(request *pb.TrackingRequest) (err error)
}

type TrackerClient interface {
	Ping() (err error)
	SetStatus(pingTime, latency int64)
	Status() (pingTime, latency int64)
	SendTrackingRequest(request *pb.TrackingRequest) (err error)
}

type trackerServerModule struct {
	pb.UnimplementedTrackerServer
	trackers map[string]TrackerClient
	mu       sync.RWMutex
}

type trackerClientEntity struct {
	id         string
	stream     pb.Tracker_SubscribeServer
	lastPinged int64
	latency    int64
}

func (m *module) InitializeTrackerServer(port int) (err error) {
	var lis net.Listener
	grpcServer := grpc.NewServer()
	if lis, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", port)); err != nil {
		return err
	}
	m.trackerServer = &trackerServerModule{trackers: make(map[string]TrackerClient)}
	// TODO: authentication
	pb.RegisterTrackerServer(grpcServer, m.trackerServer)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("[TrackerServer] failed starting server. %v", err)
		}
	}()
	return
}

// implements pb.UnimplementedTrackerServer
func (s *trackerServerModule) Subscribe(request *pb.SubscriptionRequest, stream pb.Tracker_SubscribeServer) (err error) {
	trackerID := request.TrackerId
	// TODO: validate client
	s.mu.Lock()
	if _, ok := s.trackers[trackerID]; ok {
		s.mu.Unlock()
		return fmt.Errorf("trackerID %s already registered", trackerID)
	}
	s.trackers[trackerID] = &trackerClientEntity{
		id:     trackerID,
		stream: stream,
	}
	s.mu.Unlock() // immediately unlock
	log.Printf("[TrackerServer] tracker %s connected. connected trackers: %d\n", trackerID, len(s.trackers))
	return s.PingTracker(trackerID) // keep alive: locking
}

// implements pb.UnimplementedTrackerServer
func (s *trackerServerModule) ReportTrackingRequest(ctx context.Context, message *pb.TrackingMessage) (*pb.Empty, error) {
	// TODO: implement
	response := message.Body.(*pb.TrackingMessage_Response).Response
	log.Printf("[TrackerServer] receive tracking response from %s: %+v", response.TrackerId, response)
	return &pb.Empty{}, nil
}

// implements pb.UnimplementedTrackerServer
func (s *trackerServerModule) Pong(ctx context.Context, message *pb.TrackingMessage) (*pb.Empty, error) {
	request := message.Body.(*pb.TrackingMessage_Request).Request
	trackerID := request.TrackerId
	latency := time.Now().UnixNano() - request.RequestedAt
	log.Printf("[TrackerServer] latency to %s: %dms", trackerID, latency/1e6)
	s.mu.RLock()
	defer s.mu.RUnlock()
	if client, ok := s.trackers[trackerID]; !ok {
		return &pb.Empty{}, fmt.Errorf("unregistered trackerID %s", request.TrackerId)
	} else {
		client.SetStatus(time.Now().Unix(), latency)
	}
	return &pb.Empty{}, nil
}

func (s *trackerServerModule) PingTracker(trackerID string) (err error) {
	var client TrackerClient
	var ok bool
	s.mu.RLock()
	if client, ok = s.trackers[trackerID]; !ok {
		return fmt.Errorf("unregistered trackerID %s", trackerID)
	}
	s.mu.RUnlock() // immediately unlock
	for {
		err = client.Ping()
		if err != nil {
			s.mu.Lock()
			defer s.mu.Unlock()
			delete(s.trackers, trackerID)
			log.Printf("[TrackerServer] %s disconnected. connected trackers: %d\n", trackerID, len(s.trackers))
			return
		}
		time.Sleep(constants.TrackerPingInterval)
	}
}

func (s *trackerServerModule) SendTrackingRequest(request *pb.TrackingRequest) (err error) {
	// TODO: example setup
	s.mu.Lock()
	defer s.mu.Unlock()
	if client, ok := s.trackers[request.TrackerId]; ok {
		return client.SendTrackingRequest(request)
	}
	return fmt.Errorf("unregistered trackerID %s", request.TrackerId)
}

func (c *trackerClientEntity) Ping() (err error) {
	return c.stream.Send(&pb.TrackingMessage{
		Code: pb.MessageType_MESSAGE_TYPE_PING,
		Body: &pb.TrackingMessage_Request{
			Request: &pb.TrackingRequest{
				TrackerId:   c.id,
				RequestedAt: time.Now().UnixNano(), // only PING uses Unix nano
			},
		},
	})
}

func (c *trackerClientEntity) SetStatus(pingTime, latency int64) {
	c.lastPinged = pingTime
	c.latency = latency
}

func (c *trackerClientEntity) Status() (pingTime, latency int64) {
	return c.lastPinged, c.latency
}

func (c *trackerClientEntity) SendTrackingRequest(request *pb.TrackingRequest) (err error) {
	return c.stream.Send(&pb.TrackingMessage{
		Code: pb.MessageType_MESSAGE_TYPE_TRACKING,
		Body: &pb.TrackingMessage_Request{
			Request: request,
		},
	})
}
