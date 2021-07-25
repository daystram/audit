package handler

import (
	"log"
	"time"

	"github.com/go-ping/ping"

	"github.com/daystram/audit/audit-tr/config"
	pb "github.com/daystram/audit/proto"
)

func (m *module) TrackPING(request *pb.TrackingRequest) pb.TrackingResponse {
	log.Printf("[Track] PING @ %s\n", request.Endpoint)
	init := time.Now()
	var err error
	var pinger *ping.Pinger
	if pinger, err = ping.NewPinger(request.Endpoint); err != nil {
		log.Println(err)
		return pb.TrackingResponse{
			ApplicationId: request.ApplicationId,
			ServiceId:     request.ServiceId,
			TrackerId:     config.AppConfig.TrackerID,
			Status:        pb.ServiceStatus_SERVICE_STATUS_UNKNOWN,
			Body:          err.Error(),
			ExecutedAt:    init.UnixNano(),
		}
	}
	pinger.Timeout = 1 * time.Second
	pinger.Interval = 200 * time.Millisecond
	if err = pinger.Run(); err != nil {
		log.Println(err)
		return pb.TrackingResponse{
			ApplicationId: request.ApplicationId,
			ServiceId:     request.ServiceId,
			TrackerId:     config.AppConfig.TrackerID,
			Status:        pb.ServiceStatus_SERVICE_STATUS_UNREACHABLE,
			Body:          err.Error(),
			ExecutedAt:    init.UnixNano(),
		}
	}
	return pb.TrackingResponse{
		ApplicationId: request.ApplicationId,
		ServiceId:     request.ServiceId,
		TrackerId:     config.AppConfig.TrackerID,
		Status:        pb.ServiceStatus_SERVICE_STATUS_UP,
		ResponseTime:  int64(pinger.Statistics().AvgRtt),
		ExecutedAt:    init.UnixNano(),
	}
}
