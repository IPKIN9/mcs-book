package db

import (
	"time"
)

type Author struct {
	AuthorID    int64  `gorm:"primary_key"`
	FirstName   string `gorm:"type:varchar(100)"`
	LastName    string `gorm:"type:varchar(20)"`
	Biography   string `gorm:"type:varchar(20)"`
	DateOfBirth time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
