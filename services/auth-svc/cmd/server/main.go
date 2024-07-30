package main

import (
	"auth-svc/pkg/db"
	"auth-svc/pkg/protos"
	"auth-svc/pkg/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db.Init()

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protos.RegisterUserServiceServer(s, &services.UserService{})
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
