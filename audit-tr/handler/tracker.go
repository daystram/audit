package handler

import (
	"context"
	"log"

	pb "github.com/daystram/audit/proto"
)

func (m *module) SubscribeTracking() (err error) {
	for {
		var message *pb.TrackingMessage
		if message, err = m.stream.Recv(); err != nil {
			return err
		}
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
	request := message.Body.(*pb.TrackingMessage_Request).Request
	var response pb.TrackingResponse
	switch request.Type {
	case pb.ServiceType_SERVICE_TYPE_HTTP:
		response = m.TrackHTTP(request)
	case pb.ServiceType_SERVICE_TYPE_TCP:
		response = m.TrackTCP(request)
	case pb.ServiceType_SERVICE_TYPE_PING:
		response = m.TrackPING(request)
	default:
		log.Printf("unsupported service type %s\n", request.Type)
	}

	_, err = m.client.ReportTrackingRequest(context.Background(), &pb.TrackingMessage{
		Code: pb.MessageType_MESSAGE_TYPE_TRACKING,
		Body: &pb.TrackingMessage_Response{
			Response: &response,
		},
	})
	return
}
