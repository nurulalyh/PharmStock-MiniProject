package models

import (
	"time"

	"gorm.io/gorm"
)

type DetailProduct struct {
	Id            int            `gorm:"type:smallint;not null" json:"id" form:"id"`
	NoBatch       string         `gorm:"primaryKey;type:varchar(25)" json:"no_batch" form:"no_batch"`
	IdProduct     int            `gorm:"type:smallint;not null" json:"id_product" form:"id_product"`
	MfDate        time.Time      `gorm:"type:date;not null" json:"mf_date" form:"mf_date"`
	ExpDate       time.Time      `gorm:"type:date;not null" json:"exp_date" form:"exp_date"`
	UnitPrice     int            `gorm:"type:smallint;not null" json:"unit_price" form:"unit_price"`
	Stock         int            `gorm:"type:smallint;not null" json:"stock" form:"stock"`
	IdDistributor int            `gorm:"type:smallint;not null" json:"id_distributor" form:"id_distributor"`
	CreatedAt     time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt     time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	Transactions  []Transaction  `gorm:"foreignKey:IdDetailProduct;references:NoBatch"`
}
