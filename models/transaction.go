package models

import (
	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	Id                string               `gorm:"primaryKey;type:varchar(10)"`
	IdEmployee        string               `gorm:"type:varchar(10);not null"`
	TotalQuantity     int                  `gorm:"type:smallint;not null"`
	TotalPrice        int                  `gorm:"type:smallint;not null"`
	Type              string               `gorm:"type:ENUM('inbound','outbound');not null"`
	CreatedAt         time.Time            `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time            `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt         gorm.DeletedAt       `gorm:"index"`
	DetailTransaction []DetailTransactions `gorm:"foreignKey:id_transaction;references:id"`
}
