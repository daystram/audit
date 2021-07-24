package handler

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/daystram/audit/audit-tr/config"
	pb "github.com/daystram/audit/proto"
)

type HandlerFunc interface {
	SubscribeTracking()
}

type module struct {
	conn   *grpc.ClientConn
	client pb.TrackerClient
	stream pb.Tracker_SubscribeClient
}

func InitializeHandler(conn *grpc.ClientConn) (handler *module, err error) {
	// Connect to audit-be
	var stream pb.Tracker_SubscribeClient
	client := pb.NewTrackerClient(conn)
	if stream, err = client.Subscribe(context.Background(), &pb.SubscriptionRequest{
		TrackerId: config.AppConfig.TrackerID,
	}); err != nil {
		log.Fatalf("[INIT] Failed dialling to audit-be. %v", err)
	}
	log.Println("[INIT] Successfully connected to audit-be")

	handler = &module{
		conn:   conn,
		client: client,
		stream: stream,
	}
	return
}
