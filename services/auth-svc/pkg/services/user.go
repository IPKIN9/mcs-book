package services

import (
	"auth-svc/pkg/db"
	"auth-svc/pkg/protos"
	"auth-svc/utils"
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(ctx context.Context, req *protos.CreateUserRequest) (*protos.User, error) {
	now := time.Now()

	hashed, err := utils.CreateHashing(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cant hash password")
	}

	user := db.User{
		Username:  req.Username,
		Password:  string(hashed),
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "error cant create data %v", err)
	}

	return &protos.User{
		UserId:    user.UserId,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: timestamppb.New(*user.CreatedAt),
		UpdatedAt: timestamppb.New(*user.UpdatedAt),
	}, nil
}
