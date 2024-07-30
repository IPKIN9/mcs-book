package main

import (
	"book-svc/pkg/db"
	"book-svc/pkg/protos"
	service "book-svc/pkg/services"
	"log"
	"net"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db.Init()
	defer func() {
		log.Println("error: ", string(debug.Stack()))
	}()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	client, _ := service.NewAuthorServiceClient()
	bookSvc := service.NewBookService(client)
	protos.RegisterBookServiceServer(s, bookSvc)

	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
