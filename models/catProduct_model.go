package models

import (
	"time"

	"gorm.io/gorm"
)

type CatProduct struct {
	Id             int             `gorm:"primaryKey;type:smallint" json:"id" form:"id"`
	Name           string          `gorm:"type:varchar(100);not null" json:"name" form:"name"`
	CreatedAt      time.Time       `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt      time.Time       `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt      gorm.DeletedAt  `gorm:"index" json:"deleted_at" form:"deleted_at"`
	DetailProducts []DetailProduct `gorm:"foreignKey:IdDistributor;references:Id"`
}
