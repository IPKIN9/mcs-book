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
	"gorm.io/gorm"
)

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetToken(ctx context.Context, req *protos.AuthRequest) (*protos.AuthResponse, error) {
	var findUser *db.User

	if err := db.DB.Table("users").Where("username = ?", req.Username).First(&findUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "user not find")
		}

		return nil, status.Errorf(codes.Internal, "cant find user")
	}

	passwordMatch := utils.CheckPasswordHash(findUser.Password, req.Password)
	if !passwordMatch {
		return nil, status.Errorf(codes.InvalidArgument, "password incorrect")
	}

	tokenString, expirationTime, err := utils.GenerateToken(req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get token")
	}

	var findToken *db.Token
	tx := db.DB.Begin()
	now := time.Now()

	if err := db.DB.Table("tokens").Where("user_id = ?", findUser.UserId).First(&findToken).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.Internal, "failed to get token")
		}
		findToken.Token = tokenString
		findToken.UserId = findUser.UserId
		findToken.Expired = expirationTime
		findToken.CreatedAt = &now
		findToken.UpdatedAt = &now
		if err := db.DB.Table("tokens").Create(&findToken).Error; err != nil {
			tx.Rollback()
			return nil, status.Errorf(codes.Internal, "failed to get token")
		}
	} else {
		findToken.Token = tokenString
		findToken.Expired = expirationTime
		findToken.UpdatedAt = &now

		if err := db.DB.Table("tokens").Where("token_id = ?", findToken.TokenId).Updates(&findToken).Error; err != nil {
			tx.Rollback()
			return nil, status.Errorf(codes.Internal, "failed to get token")
		}
	}

	tx.Commit()
	return &protos.AuthResponse{
		Token:   tokenString,
		Expired: timestamppb.New(expirationTime),
	}, nil
}

func (s *UserService) ValidateToken(ctx context.Context, req *protos.ValidateTokenRequest) (*protos.User, error) {
	var user *db.User
	claim, err := utils.VerifyToken(req.Token)

	if err != nil {
		return nil, err
	}

	if err := db.DB.Table("users").Where("username = ?", claim.Username).Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "cant find user")
		}

		return nil, status.Errorf(codes.Internal, "invalid token : %v", err)
	}

	return &protos.User{
		UserId:    user.UserId,
		Username:  claim.Username,
		FirstName: user.Username,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: timestamppb.New(*user.CreatedAt),
		UpdatedAt: timestamppb.New(*user.UpdatedAt),
	}, nil
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
