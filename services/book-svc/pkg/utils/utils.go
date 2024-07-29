package utils

import (
	"book-svc/pkg/db"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func UniqueIsbn(ctx context.Context, newIsbn string, bookId *int64) error {
	var existingBook []db.Book
	var err error

	err = db.DB.Where("isbn = ?", newIsbn).Find(&existingBook).Error
	same := 0

	if err == nil {
		for _, book := range existingBook {
			if bookId != nil && bookId != bookId && book.ISBN == newIsbn {
				same += 1
			}

			if bookId == nil && book.ISBN == newIsbn {
				same += 1
			}
		}
	} else if err != gorm.ErrRecordNotFound {
		return status.Errorf(codes.Internal, "failed to check ISBN: %v", err)
	}

	if same >= 1 {
		return status.Errorf(codes.AlreadyExists, "book with ISBN %s already exists", newIsbn)
	}

	return nil
}
