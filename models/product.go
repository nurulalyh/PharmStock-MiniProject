package models

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	Id                string               `gorm:"primaryKey;type:varchar(10)"`
	Name              string               `gorm:"type:varchar(100);not null"`
	Photo             string               `gorm:"type:text;not null"`
	IdCatProduct      string               `gorm:"type:varchar(10);not null"`
	MfDate            time.Time            `gorm:"type:date;not null"`
	ExpDate           time.Time            `gorm:"type:date;not null"`
	BatchNumber       int                  `gorm:"type:smallint;not null"`
	UnitPrice         int                  `gorm:"type:smallint;not null"`
	Stock             int                  `gorm:"type:smallint;not null"`
	Description       string               `gorm:"type:text;not null"`
	IdDistributor     string               `gorm:"type:varchar(10);not null"`
	CreatedAt         time.Time            `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time            `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP"`
	DeletedAt         gorm.DeletedAt       `gorm:"index"`
	DetailTransaction []DetailTransactions `gorm:"foreignKey:id_product;references:Id"`
}
