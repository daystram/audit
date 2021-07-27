package handler

import (
	"log"
	"net"
	"time"

	"github.com/daystram/audit/audit-tr/config"
	pb "github.com/daystram/audit/proto"
)

func (m *module) TrackTCP(request *pb.TrackingRequest) pb.TrackingResponse {
	log.Printf("[Track] TCP  @ %s\n", request.Endpoint)
	var init, start, end time.Time
	init = time.Now()
	start = time.Now()
	var err error
	var conn net.Conn
	if conn, err = net.DialTimeout("tcp", request.Endpoint, 1*time.Second); err != nil {
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
	conn.Close()
	end = time.Now()
	return pb.TrackingResponse{
		ApplicationId: request.ApplicationId,
		ServiceId:     request.ServiceId,
		TrackerId:     config.AppConfig.TrackerID,
		Status:        pb.ServiceStatus_SERVICE_STATUS_UP,
		ResponseTime:  end.Sub(start).Nanoseconds(),
		ExecutedAt:    init.UnixNano(),
	}
}
