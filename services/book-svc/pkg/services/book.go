package service

import (
	"book-svc/pkg/db"
	"book-svc/pkg/protos"
	"book-svc/pkg/utils"
	"context"
	"errors"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func (s *BookService) GetBook(ctx context.Context, req *protos.GetBookRequest) (*protos.GetBookResponse, error) {
	var book db.Book
	if err := db.DB.First(&book, req.BookId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Book with ID %d not found", req.BookId)
		}
		return nil, status.Errorf(codes.Internal, "Database error: %v", err)
	}

	return &protos.GetBookResponse{
		Book: &protos.Book{
			BookId:        book.BookID,
			Title:         book.Title,
			Isbn:          book.ISBN,
			AuthorId:      book.AuthorID,
			CategoryId:    book.CategoryID,
			PublishedDate: timestamppb.New(book.PublishedDate),
			Description:   book.Description,
			CreatedAt:     timestamppb.New(*book.CreatedAt),
			UpdatedAt:     timestamppb.New(*book.UpdatedAt),
		},
	}, nil
}

func (s *BookService) GetAllBooks(ctx context.Context, req *protos.GetAllBookRequest) (*protos.GetAllBooksResponse, error) {
	var books []db.Book

	tb := db.DB.Table("books")
	if req.Search != nil && len(req.Search.Value) >= 1 {
		searchValue := "%" + req.Search.Value + "%"
		tb = tb.Where("title ILIKE ?", searchValue)
	}
	if err := tb.Limit(int(req.Limit)).Offset(int(req.Page-1) * int(req.Limit)).Find(&books).Error; err != nil {
		return nil, err
	}

	var protoBooks []*protos.Book
	for _, book := range books {
		protoBooks = append(protoBooks, &protos.Book{
			BookId:        book.BookID,
			Title:         book.Title,
			Isbn:          book.ISBN,
			AuthorId:      book.AuthorID,
			CategoryId:    book.CategoryID,
			PublishedDate: timestamppb.New(book.PublishedDate),
			Description:   book.Description,
			CreatedAt:     timestamppb.New(*book.CreatedAt),
			UpdatedAt:     timestamppb.New(*book.UpdatedAt),
		})
	}

	return &protos.GetAllBooksResponse{Books: protoBooks}, nil
}

func (s *BookService) CreateBook(ctx context.Context, req *protos.CreateBookRequest) (*protos.BookResponse, error) {
	now := time.Now()
	book := db.Book{
		Title:         req.Title,
		ISBN:          req.Isbn,
		AuthorID:      req.AuthorId,
		CategoryID:    req.CategoryId,
		PublishedDate: time.Unix(req.PublishedDate.Seconds, int64(req.PublishedDate.Nanos)),
		Description:   req.Description,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}

	if _err := utils.UniqueIsbn(ctx, book.ISBN, nil); _err != nil {
		return nil, _err
	}

	if err := db.DB.Create(&book).Error; err != nil {
		return nil, err
	}

	return &protos.BookResponse{
		Book: &protos.Book{
			BookId:        book.BookID,
			Title:         book.Title,
			Isbn:          book.ISBN,
			AuthorId:      book.AuthorID,
			CategoryId:    book.CategoryID,
			PublishedDate: timestamppb.New(book.PublishedDate),
			Description:   book.Description,
			CreatedAt:     timestamppb.New(*book.CreatedAt),
			UpdatedAt:     timestamppb.New(*book.UpdatedAt),
		},
	}, nil
}

func (s *BookService) UpdateBook(ctx context.Context, req *protos.UpdateBookRequest) (*protos.BookResponse, error) {
	var book db.Book

	if err := db.DB.First(&book, req.BookId).Error; err != nil {
		log.Printf("eror ini: %v", err)
		return nil, err
	}
	now := time.Now()
	book.Title = req.Book.Title
	book.ISBN = req.Book.Isbn
	book.AuthorID = req.Book.AuthorId
	book.CategoryID = req.Book.CategoryId
	book.PublishedDate = time.Unix(req.Book.PublishedDate.Seconds, int64(req.Book.PublishedDate.Nanos))
	book.Description = req.Book.Description
	book.UpdatedAt = &now

	if _err := utils.UniqueIsbn(ctx, book.ISBN, &req.BookId); _err != nil {
		return nil, _err
	}

	if err := db.DB.Save(&book).Error; err != nil {
		return nil, err
	}

	return &protos.BookResponse{
		Book: &protos.Book{
			BookId:        book.BookID,
			Title:         book.Title,
			Isbn:          book.ISBN,
			AuthorId:      book.AuthorID,
			CategoryId:    book.CategoryID,
			PublishedDate: timestamppb.New(book.PublishedDate),
			Description:   book.Description,
			CreatedAt:     timestamppb.New(*book.CreatedAt),
			UpdatedAt:     timestamppb.New(*book.UpdatedAt),
		},
	}, nil
}

func (s *BookService) DeleteBook(ctx context.Context, req *protos.DeleteBookRequest) (*protos.Empty, error) {
	if err := db.DB.Delete(&db.Book{}, req.BookId).Error; err != nil {
		return nil, err
	}

	return &protos.Empty{}, nil
}
