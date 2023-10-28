package response

import (
	"time"

	"gorm.io/gorm"
)

type TransactionResponse struct {
	Id            string         `json:"id" form:"id"`
	IdEmployee    string         `json:"id_employee" form:"id_employee"`
	TotalQuantity int            `json:"total_quantity" form:"total_quantity"`
	TotalPrice    int            `json:"total_price" form:"total_price"`
	Type          string         `json:"type" form:"type"`
	CreatedAt     time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
}
