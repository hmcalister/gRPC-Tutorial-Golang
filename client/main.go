package main

import (
	"context"
	pb "hmcalister/grpcTutorial/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	ADDRESS = "localhost:50051"
)

type ActivityTask struct {
	Name      string
	Important bool
}

func main() {
	conn, err := grpc.Dial(ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("gRPC could not dial: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewActivityServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	activites := []ActivityTask{
		{Name: "Activity 1", Important: false},
		{Name: "Activity 2", Important: true},
		{Name: "Activity 3", Important: false},
		{Name: "Activity 4", Important: true},
		{Name: "Activity 5", Important: false},
	}

	for _, activity := range activites {
		response, err := grpcClient.CreateActivity(ctx, &pb.NewActivity{
			Name:      activity.Name,
			Important: activity.Important,
		})
		if err != nil {
			log.Fatalf("Could not create activity %#v: Err %v", activity, err)
		}

		log.Printf(`
			ID: %s
			Name: %s
			Important %v
		`, response.GetId(), response.GetName(), response.GetImportant())
	}
}
