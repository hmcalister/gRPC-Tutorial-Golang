package main

import (
	"context"
	pb "hmcalister/grpcTutorial/proto"
	"log"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	PORT = ":50051"
)

type ActivityServer struct {
	pb.UnimplementedActivityServiceServer
}

func (s *ActivityServer) CreateActivity(ctx context.Context, in *pb.NewActivity) (*pb.Activity, error) {
	log.Printf("Received: %v", in.GetName())

	activity := &pb.Activity{
		Name:      in.GetName(),
		Important: in.GetImportant(),
		Id:        uuid.New().String(),
	}

	return activity, nil
}

func main() {
	listenConn, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to listen with error: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterActivityServiceServer(grpcServer, &ActivityServer{})
	log.Printf("gRPC server listening on %v", listenConn.Addr())

	if err := grpcServer.Serve(listenConn); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
