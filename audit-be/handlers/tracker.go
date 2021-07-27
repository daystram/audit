package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/daystram/audit/audit-be/constants"
	"github.com/daystram/audit/audit-be/models"
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
	handlers   *module // injected handler module reference
	trackers   map[string]TrackerClient
	trackerIDs []string
	lastUsed   int
	mu         sync.RWMutex
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
	if lis, err = net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port)); err != nil {
		return err
	}
	log.Printf("[TrackerServer] starting server on 0.0.0.0:%d\n", port)
	m.trackerServer = &trackerServerModule{handlers: m, trackers: make(map[string]TrackerClient), trackerIDs: make([]string, 0)}
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
	s.trackerIDs = append(s.trackerIDs, trackerID)
	s.mu.Unlock() // immediately unlock
	log.Printf("[TrackerServer] tracker %s connected. connected trackers: %d\n", trackerID, len(s.trackers))
	return s.PingTracker(trackerID) // keep alive: locking
}

// implements pb.UnimplementedTrackerServer
func (s *trackerServerModule) ReportTrackingRequest(ctx context.Context, message *pb.TrackingMessage) (resp *pb.Empty, err error) {
	response := message.Body.(*pb.TrackingMessage_Response).Response
	var application models.Application
	if application, err = s.handlers.db.applicationOrmer.GetOneByID(response.ApplicationId); err != nil {
		return &pb.Empty{}, nil
	}
	var service models.Service
	if service, err = s.handlers.db.serviceOrmer.GetOneByIDAndApplicationID(response.ServiceId, response.ApplicationId); err != nil {
		return &pb.Empty{}, nil
	}
	switch response.Status {
	case pb.ServiceStatus_SERVICE_STATUS_UP:
		switch service.Type {
		case constants.ServiceTypeHTTP:
			var code int
			if code, err = strconv.Atoi(response.Body); err != nil {
				return &pb.Empty{}, fmt.Errorf("invalid body for HTTP service")
			}
			log.Printf("[TrackerServer] HTTP %s/%s -> latency:%dms, code:%d\n", application.Name, service.Name, response.ResponseTime/1e6, code)
			// TODO
		case constants.ServiceTypeTCP:
			log.Printf("[TrackerServer] TCP  %s/%s -> latency:%dms\n", application.Name, service.Name, response.ResponseTime/1e6)
			// TODO
		case constants.ServiceTypePING:
			log.Printf("[TrackerServer] PING %s/%s -> latency:%dms\n", application.Name, service.Name, response.ResponseTime/1e6)
			// TODO
		}
		if err = s.handlers.influx.reportOrmer.Insert(models.Report{
			ApplicationID: response.ApplicationId,
			ServiceID:     response.ServiceId,
			Latency:       response.ResponseTime,
			Timestamp:     response.ExecutedAt,
		}); err != nil {
			log.Println(err)
		}
	case pb.ServiceStatus_SERVICE_STATUS_UNREACHABLE:
		log.Printf("[TrackerServer] DOWN %s/%s @ trackerID:%s -> unreachable: %s\n", application.Name, service.Name, response.TrackerId, response.Body)
	default:
		log.Printf("[TrackerServer] ERR  %s/%s @ trackerID:%s -> err: %s\n", application.Name, service.Name, response.TrackerId, response.Body)
	}
	return &pb.Empty{}, nil
}

// implements pb.UnimplementedTrackerServer
func (s *trackerServerModule) Pong(ctx context.Context, message *pb.TrackingMessage) (*pb.Empty, error) {
	request := message.Body.(*pb.TrackingMessage_Request).Request
	trackerID := request.TrackerId
	latency := time.Now().UnixNano() - request.RequestedAt
	log.Printf("[TrackerServer] PING-TR trackerID:%s -> latency:%dms", trackerID, latency/1e6)
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
			for i, ID := range s.trackerIDs {
				if trackerID == ID {
					s.trackerIDs = append(s.trackerIDs[:i], s.trackerIDs[i+1:]...)
					break
				}
			}
			if len(s.trackers) != len(s.trackerIDs) {
				log.Fatalln("trackers desync")
				return
			}
			log.Printf("[TrackerServer] trackerID:%s disconnected. connected trackers: %d\n", trackerID, len(s.trackers))
			return
		}
		time.Sleep(constants.TrackerPingInterval)
	}
}

func (s *trackerServerModule) SendTrackingRequest(request *pb.TrackingRequest) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// round-robin
	if len(s.trackers) == 0 {
		return errors.New("no registered trackers")
	}
	s.lastUsed = (s.lastUsed + 1) % len(s.trackers)
	request.TrackerId = s.trackerIDs[s.lastUsed]
	return s.trackers[request.TrackerId].SendTrackingRequest(request)
}

func (c *trackerClientEntity) Ping() (err error) {
	return c.stream.Send(&pb.TrackingMessage{
		Code: pb.MessageType_MESSAGE_TYPE_PING,
		Body: &pb.TrackingMessage_Request{
			Request: &pb.TrackingRequest{
				TrackerId:   c.id,
				RequestedAt: time.Now().UnixNano(),
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
