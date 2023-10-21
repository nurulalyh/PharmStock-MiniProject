package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id           int            `gorm:"primaryKey;type:smallint" json:"id" form:"id"`
	Name         string         `gorm:"type:varchar(100);not null" json:"name" form:"name"`
	Photo        string         `gorm:"type:text;not null" json:"photo" form:"photo"`
	Description  string         `gorm:"type:text;not null" json:"description" form:"description"`
	TotStock     int            `gorm:"type:smallint;not null" json:"tot_stock" form:"tot_stock"`
	IdCatProduct int            `gorm:"type:smallint;not null" json:"id_cat_product" form:"id_cat_product"`
	CreatedAt    time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt    time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	DetailProducts []DetailProduct `gorm:"foreignKey:IdProduct;references:Id"`
}
