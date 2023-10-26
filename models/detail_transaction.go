package models

import (
	"time"

	"gorm.io/gorm"
)

type DetailTransactions struct {
	Id            string         `gorm:"primaryKey;type:varchar(10)"`
	IdTransaction string         `gorm:"type:varchar(10);not null"`
	IdProduct     string         `gorm:"type:varchar(10);not null"`
	Quantity      int            `gorm:"type:smallint;not null"`
	Price         int            `gorm:"type:smallint;not null"`
	CreatedAt     time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
