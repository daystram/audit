package handler

import (
	"context"
	"log"
	"time"

	"github.com/daystram/audit/audit-tr/config"
	pb "github.com/daystram/audit/proto"
)

func (m *module) SubscribeTracking() (err error) {
	for {
		var message *pb.TrackingMessage
		if message, err = m.stream.Recv(); err != nil {
			return err
		}
		log.Println(message)
		switch message.Code {
		case pb.MessageType_MESSAGE_TYPE_PING:
			if err = m.ReplyPing(message); err != nil {
				return
			}
		case pb.MessageType_MESSAGE_TYPE_TRACKING:
			if err = m.ExecuteTracking(message); err != nil {
				return
			}
		default:
			log.Printf("unsupported message type %s\n", message.Code.String())
		}
	}
}

func (m *module) ReplyPing(message *pb.TrackingMessage) (err error) {
	_, err = m.client.Pong(context.Background(), message)
	return
}

func (m *module) ExecuteTracking(message *pb.TrackingMessage) (err error) {
	// TODO: implement
	request := message.Body.(*pb.TrackingMessage_Request)
	time.Sleep(2 * time.Second) //simulate
	_, err = m.client.ReportTrackingRequest(context.Background(), &pb.TrackingMessage{
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
	return
}
