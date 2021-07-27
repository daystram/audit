package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/daystram/audit/audit-tr/config"
	pb "github.com/daystram/audit/proto"
)

func (m *module) TrackHTTP(request *pb.TrackingRequest) pb.TrackingResponse {
	log.Printf("[Track] HTTP @ %s\n", request.Endpoint)
	var init, start, end time.Time
	init = time.Now()
	var err error
	var httpRequest *http.Request
	if httpRequest, err = http.NewRequest(http.MethodGet, request.Endpoint, nil); err != nil {
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
	client := http.Transport{
		IdleConnTimeout:     1 * time.Second,
		TLSHandshakeTimeout: 1 * time.Second,
		DisableKeepAlives:   true,
	}
	var httpResponse *http.Response
	if httpResponse, err = client.RoundTrip(
		httpRequest.WithContext(httptrace.WithClientTrace(httpRequest.Context(), &httptrace.ClientTrace{
			ConnectStart: func(_, _ string) {
				start = time.Now()
			},
		})),
	); err != nil {
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
	end = time.Now()
	return pb.TrackingResponse{
		ApplicationId: request.ApplicationId,
		ServiceId:     request.ServiceId,
		TrackerId:     config.AppConfig.TrackerID,
		Status:        pb.ServiceStatus_SERVICE_STATUS_UP,
		Body:          fmt.Sprint(httpResponse.StatusCode),
		ResponseTime:  end.Sub(start).Nanoseconds(),
		ExecutedAt:    init.UnixNano(),
	}
}
