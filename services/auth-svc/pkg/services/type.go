package services

import (
	"auth-svc/pkg/protos"
	"context"
)

type userService interface {
	CreateUser(ctx context.Context, req *protos.CreateUserRequest) (*protos.User, error)
}

type UserService struct {
	protos.UnimplementedUserServiceServer
	service userService
}
