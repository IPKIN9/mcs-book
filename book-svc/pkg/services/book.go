package service

import (
	"book-svc/pkg/db"
	"book-svc/pkg/protos"
	"context"
	"errors"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type BookService struct {
	protos.UnimplementedBookServiceServer
}

func uniqueIsbn(ctx context.Context, newIsbn string, bookId *int64) error {
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
			CreatedAt:     timestamppb.New(book.CreatedAt),
			UpdatedAt:     timestamppb.New(book.UpdatedAt),
		},
	}, nil
}

func (s *BookService) GetAllBooks(ctx context.Context, req *protos.Empty) (*protos.GetAllBooksResponse, error) {
	var books []db.Book
	if err := db.DB.Find(&books).Error; err != nil {
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
			CreatedAt:     timestamppb.New(book.CreatedAt),
			UpdatedAt:     timestamppb.New(book.UpdatedAt),
		})
	}

	return &protos.GetAllBooksResponse{Books: protoBooks}, nil
}

func (s *BookService) CreateBook(ctx context.Context, req *protos.CreateBookRequest) (*protos.BookResponse, error) {
	book := db.Book{
		Title:         req.Title,
		ISBN:          req.Isbn,
		AuthorID:      req.AuthorId,
		CategoryID:    req.CategoryId,
		PublishedDate: time.Unix(req.PublishedDate.Seconds, int64(req.PublishedDate.Nanos)),
		Description:   req.Description,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if _err := uniqueIsbn(ctx, book.ISBN, nil); _err != nil {
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
			CreatedAt:     timestamppb.New(book.CreatedAt),
			UpdatedAt:     timestamppb.New(book.UpdatedAt),
		},
	}, nil
}

func (s *BookService) UpdateBook(ctx context.Context, req *protos.UpdateBookRequest) (*protos.BookResponse, error) {
	var book db.Book

	if err := db.DB.First(&book, req.BookId).Error; err != nil {
		log.Printf("eror ini: %v", err)
		return nil, err
	}

	book.Title = req.Book.Title
	book.ISBN = req.Book.Isbn
	book.AuthorID = req.Book.AuthorId
	book.CategoryID = req.Book.CategoryId
	book.PublishedDate = time.Unix(req.Book.PublishedDate.Seconds, int64(req.Book.PublishedDate.Nanos))
	book.Description = req.Book.Description
	book.UpdatedAt = time.Now()

	if _err := uniqueIsbn(ctx, book.ISBN, &req.BookId); _err != nil {
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
			CreatedAt:     timestamppb.New(book.CreatedAt),
			UpdatedAt:     timestamppb.New(book.UpdatedAt),
		},
	}, nil
}

func (s *BookService) DeleteBook(ctx context.Context, req *protos.DeleteBookRequest) (*protos.Empty, error) {
	if err := db.DB.Delete(&db.Book{}, req.BookId).Error; err != nil {
		return nil, err
	}

	return &protos.Empty{}, nil
}
