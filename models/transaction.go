package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id              int            `gorm:"primaryKey;type:smallint" json:"id" form:"id"`
	IdUser          int            `gorm:"type:smallint;not null" json:"id_user" form:"id_user"`
	IdDetailProduct string         `gorm:"type:varchar(25);not null" json:"id_detail_product" form:"id_detail_product"`
	Quantity        int            `gorm:"type:smallint;not null" json:"quantity" form:"quantity"`
	Price           int            `gorm:"type:smallint;not null" json:"price" form:"price"`
	Type            string         `gorm:"type:varchar(25);not null" json:"type" form:"type"`
	CreatedAt       time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt       time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
}
