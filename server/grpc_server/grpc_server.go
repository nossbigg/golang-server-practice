package grpc_server

import (
	context "context"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type server struct {
	UnimplementedUserServer
}

func (s *server) GetUser(ctx context.Context, in *EmptyMessage) (*GetUserResponse, error) {
	result := &GetUserResponse{Name: "jeff"}
	return result, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterUserServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
