package service

import "book-svc/pkg/protos"

type BookService struct {
	protos.UnimplementedBookServiceServer
	authorClient *AuthorServiceClient
}
