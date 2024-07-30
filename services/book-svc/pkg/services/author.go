package service

import (
	"book-svc/pkg/author/protos"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthorServiceClient struct {
	client protos.AuthorServiceClient
}

func NewAuthorServiceClient() (*AuthorServiceClient, error) {
	conn, err := grpc.NewClient(
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Printf("Error creating gRPC client: %v", err)
		return nil, err
	}

	client := protos.NewAuthorServiceClient(conn)
	return &AuthorServiceClient{client: client}, nil
}

func (c *AuthorServiceClient) GetAuthor(ctx context.Context) (*protos.Author, error) {
	resp, err := c.client.GetAuthor(ctx, &protos.GetAuthorRequest{AuthorId: 2})
	return resp, err
}
