package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/herochi/orbi/service-b/internal/application"
	pb "github.com/herochi/orbi/service-b/user"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	notifyService *application.NotifyService
}

func NewGRPCServer(service *application.NotifyService) *server {
	return &server{notifyService: service}
}

func (s *server) NotifyUser(ctx context.Context, req *pb.Request) (*pb.NotifyResponse, error) {
	message, err := s.notifyService.NotifyUser(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("error notifying user: %w", err)
	}

	return &pb.NotifyResponse{Message: message}, nil
}

func StartGRPCServer(service *application.NotifyService) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error starting gRPC server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, NewGRPCServer(service))

	fmt.Println("gRPC server listening on :50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error serving: %v", err)
	}
}
