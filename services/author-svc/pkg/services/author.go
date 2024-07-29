package service

import (
	"author-svc/pkg/db"
	"author-svc/pkg/protos"
	"author-svc/utils"
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func (s *AuthorService) GetAuthor(ctx context.Context, req *protos.GetAuthorRequest) (*protos.Author, error) {
	var author db.Author

	if err := db.DB.Table("authors").Where("author_id = ?", req.AuthorId).First(&author).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.Internal, "failed to find author")
		}

		return nil, status.Errorf(codes.NotFound, "author with id %v not found", req.AuthorId)
	}

	return &protos.Author{
		AuthorId:    author.AuthorID,
		FirstName:   author.FirstName,
		LastName:    author.LastName,
		Biography:   author.Biography,
		DateOfBirth: timestamppb.New(author.DateOfBirth),
		CreatedAt:   timestamppb.New(*author.CreatedAt),
		UpdatedAt:   timestamppb.New(*author.UpdatedAt),
	}, nil
}

func (s *AuthorService) GetAllAuthor(ctx context.Context, req *protos.GetAllAuthorRequest) (*protos.GetAllAuthorResponse, error) {
	var authors []db.Author

	tb := db.DB.Table("authors")
	if req.Search != nil && len(req.Search.Value) >= 1 {
		searchValue := "%" + req.Search.Value + "%"
		tb = tb.Where("first_name ILIKE Or last_name ILIKE ? ?", searchValue, searchValue)
	}

	if err := tb.Limit(int(req.Limit)).Offset(utils.GetPage(int(req.Offset), int(req.Limit))).Find(&authors).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get author")
	}

	var authorResult []*protos.Author
	for _, author := range authors {
		authorResult = append(authorResult, &protos.Author{
			AuthorId:    author.AuthorID,
			FirstName:   author.FirstName,
			LastName:    author.LastName,
			Biography:   author.Biography,
			DateOfBirth: timestamppb.New(author.DateOfBirth),
			CreatedAt:   timestamppb.New(*author.CreatedAt),
			UpdatedAt:   timestamppb.New(*author.UpdatedAt),
		})
	}

	return &protos.GetAllAuthorResponse{
		Authors: authorResult,
	}, nil
}

func (s *AuthorService) UpdateAuthor(ctx context.Context, req *protos.UpdateAuthorRequest) (*protos.Author, error) {
	var author db.Author

	if err := db.DB.Table("authors").Where("author_id = ?", req.AuthorId).First(&author).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.Internal, "failed to find author")
		}

		return nil, status.Errorf(codes.NotFound, "author with id %v not found", req.AuthorId)
	}

	now := time.Now()
	author.Biography = req.Authors.Biography
	author.FirstName = req.Authors.FirstName
	author.LastName = req.Authors.LastName
	author.DateOfBirth = time.Unix(req.Authors.DateOfBirth.Seconds, int64(req.Authors.DateOfBirth.Nanos))
	author.UpdatedAt = &now

	tx := db.DB.Begin()

	if err := tx.Model(&db.Author{}).Table("authors").Where("author_id = ?", req.AuthorId).Updates(author).Error; err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "failed to find author")
	}
	tx.Commit()

	return &protos.Author{
		AuthorId:    author.AuthorID,
		FirstName:   author.FirstName,
		LastName:    author.LastName,
		Biography:   author.Biography,
		DateOfBirth: timestamppb.New(author.DateOfBirth),
		CreatedAt:   timestamppb.New(*author.CreatedAt),
		UpdatedAt:   timestamppb.New(*author.UpdatedAt),
	}, nil
}

func (s *AuthorService) CreateAuthor(ctx context.Context, req *protos.CreateAuthorRequest) (*protos.Author, error) {
	now := time.Now()
	author := &db.Author{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Biography:   req.Biography,
		DateOfBirth: time.Unix(req.DateOfBirth.Seconds, int64(req.DateOfBirth.Nanos)),
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	if err := db.DB.Table("authors").Create(&author).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create author")
	}

	return &protos.Author{
		AuthorId:    author.AuthorID,
		FirstName:   author.FirstName,
		LastName:    author.LastName,
		Biography:   author.Biography,
		DateOfBirth: timestamppb.New(author.DateOfBirth),
		CreatedAt:   timestamppb.New(*author.CreatedAt),
		UpdatedAt:   timestamppb.New(*author.UpdatedAt),
	}, nil
}

func (s *AuthorService) DeleteAuthor(ctx context.Context, req *protos.DeleteAuthorRequest) (*protos.Empty, error) {
	var author db.Author

	if err := db.DB.Table("authors").Where("author_id = ?", req.AuthorId).First(&author).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.Internal, "failed to delete author")
		}

		return nil, status.Errorf(codes.NotFound, "author with id %v not found", req.AuthorId)
	}

	if err := db.DB.Delete(&db.Author{}, req.AuthorId).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete author")
	}

	return &protos.Empty{}, nil
}
