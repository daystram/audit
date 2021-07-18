package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/daystram/audit/proto"
	"google.golang.org/grpc"
)

func main() {
	trackerID := flag.String("id", "", "")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8855", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[INIT] failed dialling to audit-be. %v", err)
	}
	defer conn.Close()

	client := pb.NewTrackerClient(conn)
	stream, err := client.Subscribe(context.Background(), &pb.SubscriptionRequest{
		TrackerId: *trackerID,
	})
	if err != nil {
		log.Fatalf("[INIT] failed dialling to audit-be. %v", err)
	}
	for {
		message, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(message)

		switch message.Code {
		case pb.MessageType_MESSAGE_TYPE_PING:
			client.Pong(context.Background(), message)
		case pb.MessageType_MESSAGE_TYPE_TRACKING:
			request := message.Body.(*pb.TrackingMessage_Request)
			time.Sleep(2 * time.Second) //simulate
			client.ReportTrackingRequest(context.Background(), &pb.TrackingMessage{
				Code: pb.MessageType_MESSAGE_TYPE_TRACKING,
				Body: &pb.TrackingMessage_Response{
					Response: &pb.TrackingResponse{
						ApplicationId: request.Request.ApplicationId,
						ServiceId:     request.Request.ServiceId,
						TrackerId:     *trackerID,
						Status:        pb.ServiceStatus_SERVICE_STATUS_UP,
						ResponseTime:  100,
						ExecutedAt:    time.Now().Unix(),
					},
				},
			})
		}
	}
}
