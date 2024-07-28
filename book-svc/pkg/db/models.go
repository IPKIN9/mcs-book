package db

import (
	"time"
)

type Book struct {
	BookID        int64  `gorm:"primary_key"`
	Title         string `gorm:"type:varchar(100)"`
	ISBN          string `gorm:"type:varchar(20)"`
	AuthorID      int64
	CategoryID    int64
	PublishedDate time.Time
	Description   string `gorm:"type:text"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Stock struct {
	StockID           int64 `gorm:"primary_key"`
	BookID            int64
	TotalQuantity     int64
	AvailableQuantity int64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
