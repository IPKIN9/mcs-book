package service

import (
	"book-svc/pkg/db"
	"book-svc/pkg/protos"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (s *BookService) BorrowingBook(ctx context.Context, req *protos.BorrowingBookRequest) (*protos.BorrowingBookResponse, error) {
	var bookStock []db.Stock

	err := db.DB.Table("stock").Where("book_id IN ?", req.BookId).Find(&bookStock).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.Internal, "failed to check ISBN: %v", err)
		}

		if err == gorm.ErrRecordNotFound {
			return &protos.BorrowingBookResponse{
				Total:   0,
				Message: "Stock is empty",
			}, nil
		}
	}

	totalCanBorrow := make([]int64, 0, 5)

	for _, book := range bookStock {
		if book.AvailableQuantity > 0 {
			totalCanBorrow = append(totalCanBorrow, book.BookID)
		}
	}

	update, err := updateStock(ctx, totalCanBorrow, req.IsBorrow)
	if err != nil {
		return &protos.BorrowingBookResponse{
			Total:   uint32(update),
			Message: err.Error(),
		}, nil
	}

	return &protos.BorrowingBookResponse{
		Total:   uint32(update),
		Message: "Success",
	}, nil
}

func (s *BookService) GetStock(ctx context.Context, req *protos.GetStockRequest) (*protos.GetStockResponse, error) {
	if req.BookId == nil && req.StockId == nil {
		return nil, status.Errorf(codes.InvalidArgument, "no params")
	}

	var bookStock db.Stock
	tb := db.DB.Table("stock")

	if req.BookId != nil {
		tb = tb.Where("book_id = ?", req.BookId.Value)
	}
	if req.StockId != nil {
		tb = tb.Where("stock_id = ?", req.StockId.Value)
	}

	err := tb.First(&bookStock).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.Internal, "failed to check ISBN: %v", err)
		}

		if err == gorm.ErrRecordNotFound {
			return &protos.GetStockResponse{
				TotalStock:   0,
				AvaibleStock: 0,
			}, nil
		}
	}

	return &protos.GetStockResponse{
		TotalStock:   int64(bookStock.TotalQuantity),
		AvaibleStock: int64(bookStock.AvailableQuantity),
	}, nil
}

// func (s *BookService) ChangeStock(ctx context.Context, req)

func updateStock(ctx context.Context, req []int64, isBorrow bool) (int32, error) {
	if len(req) < 1 {
		return 0, errors.New("avaible quantity not enough")
	}

	var bookStocks []db.Stock

	err := db.DB.Table("stock").Where("book_id = ?", req).Find(&bookStocks).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return 0, status.Errorf(codes.Internal, "failed to check ISBN: %v", err)
		}
	}

	if len(bookStocks) < 1 {
		return 0, status.Errorf(codes.NotFound, "failed to check ISBN: %v", err)
	}

	tx := db.DB.Begin()

	for _, bookStock := range bookStocks {
		if isBorrow {
			bookStock.AvailableQuantity -= 1
		} else if !isBorrow && bookStock.AvailableQuantity != bookStock.TotalQuantity {
			bookStock.AvailableQuantity += 1
		} else if !isBorrow && bookStock.AvailableQuantity >= bookStock.TotalQuantity {
			return 0, errors.New("maximum stock")
		}

		if err := tx.Model(&db.Stock{}).Table("stock").Where("book_id = ?", req).Updates(bookStock).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	tx.Commit()

	return int32(len(bookStocks)), nil
}
