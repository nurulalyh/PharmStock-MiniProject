package models

import (
	"time"

	"gorm.io/gorm"
)

type ReqProduct struct {
	Id          int            `gorm:"primaryKey;type:smallint" json:"id" form:"id"`
	ProductName string         `gorm:"type:varchar(100);not null" json:"product_name" form:"product_name"`
	IdUser      int            `gorm:"type:smallint;not null" json:"id_user" form:"id_user"`
	Quantity    int            `gorm:"type:smallint;not null" json:"quantity" form:"quantity"`
	Note        string         `gorm:"type:varchar(255);not null" json:"note" form:"note"`
	StatusReq   string         `gorm:"type:varchar(25);not null" json:"status_req" form:"status_req"`
	CreatedAt   time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at" form:"created_at"`
	UpdatedAt   time.Time      `gorm:"type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at" form:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
}
