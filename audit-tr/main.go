package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/daystram/audit/proto"
	"google.golang.org/grpc"
)

func main() {
	trackerID := flag.String("id", "", "")
	flag.Parse()

	conn, err := grpc.Dial("localhost:5555", grpc.WithInsecure())
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
		request, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(request)
		// d, _ := time.ParseDuration("1s")
		// time.Sleep(d)
	}
}
