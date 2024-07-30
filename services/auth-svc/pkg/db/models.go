package db

import (
	"time"
)

type Borrowing struct {
	BorrowingId int64 `gorm:"primary_key"`
	UserId      int64
	BookId      int64
	BorrowDate  time.Time
	DueDate     *time.Time
	ReturnDate  *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type User struct {
	UserId    int64  `gorm:"primary_key"`
	Username  string `gorm:"type:varchar(255)"`
	Password  string `gorm:"type:varchar(255)"`
	Email     string `gorm:"type:varchar(20)"`
	FirstName string `gorm:"type:varchar(100)"`
	LastName  string `gorm:"type:varchar(20)"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
