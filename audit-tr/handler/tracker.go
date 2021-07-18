package handler

import (
	"context"
	"log"
	"time"

	"github.com/daystram/audit/audit-tr/config"
	pb "github.com/daystram/audit/proto"
)

func (m *module) SubscribeTracking() {
	for {
		message, err := m.stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(message)
		switch message.Code {
		case pb.MessageType_MESSAGE_TYPE_PING:
			m.ReplyPing(message)
		case pb.MessageType_MESSAGE_TYPE_TRACKING:
			m.ExecuteTracking(message)
		default:
			log.Printf("unsupported message type %s\n", message.Code.String())
		}
	}
}

func (m *module) ReplyPing(message *pb.TrackingMessage) {
	m.client.Pong(context.Background(), message)
}

func (m *module) ExecuteTracking(message *pb.TrackingMessage) {
	// TODO: implement
	request := message.Body.(*pb.TrackingMessage_Request)
	time.Sleep(2 * time.Second) //simulate
	m.client.ReportTrackingRequest(context.Background(), &pb.TrackingMessage{
		Code: pb.MessageType_MESSAGE_TYPE_TRACKING,
		Body: &pb.TrackingMessage_Response{
			Response: &pb.TrackingResponse{
				ApplicationId: request.Request.ApplicationId,
				ServiceId:     request.Request.ServiceId,
				TrackerId:     config.AppConfig.TrackerID,
				Status:        pb.ServiceStatus_SERVICE_STATUS_UP,
				ResponseTime:  100,
				ExecutedAt:    time.Now().Unix(),
			},
		},
	})
}
